
import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, Router } from '@angular/router';
import { TakequizService } from './takequiz.service';
import { RouterModule } from '@angular/router';


interface Question {
  question: string;
  options: { option_id: number; value: string }[];
  correct_option: number;
}

interface Topic {
  quiz_topic: string;
  questions: Question[];
}

@Component({
  selector: 'app-quiz',
  templateUrl: './takequiz.component.html',
  styleUrls: ['./takequiz.component.css'],
  imports: [CommonModule, RouterModule]
})
export class TakeQuizComponent implements OnInit {
  quiz_topic_id :string |null= ''; 
  category: string = '';
  subcategory: string = '';
  quizTopic: string = '';
  quizTopicsList: any[] = [];
  quizData: Topic | null = null;

  currentQuestionIndex = 0;
  selectedAnswer: number | null = null;
  score = 0;
  showModal = false;
  quizStarted = false;
  loading = false;

  remainingTime: number = 100; // Example starting time: 1 minutes (300 seconds)
  showWarning: boolean = false; 

  constructor(private route: ActivatedRoute, private router: Router, private takequizService: TakequizService) {}

  ngOnInit(): void {
    
    this.quiz_topic_id= sessionStorage.getItem('currentQuizId');
    // Retrieve route parameters
    this.route.params.subscribe(params => {
      this.category = params['category'];
      this.subcategory = params['subcategory'];
      this.quizTopic = decodeURIComponent(params['quizTopic']);
      console.log(this.quizTopic);
      
      this.loadQuizData();
      this.loadQuizTopics();
    });
    
  }

  loadQuizTopics(): void {
    this.takequizService.getQuizTopics(this.category,this.subcategory).subscribe({
      next: (topics) => {
        console.log("Topics received:", topics);
        this.quizTopicsList = topics;
      },
      error: (err) => {
        console.error('Failed to load quiz topics list', err);
      }
    });
  }

  loadQuizData() {
    this.loading = true;
    this.takequizService.getQuizData(this.quizTopic).subscribe(
      data => {
        console.log("Fetched API Data:", data);

        if (!data) {
          console.error("No quiz data found for topic:", this.quizTopic);
          this.loading = false;
          return;
        }

        
        this.quizData = data;
        this.loading = false;
      },
      error => {
        console.error("Error fetching quiz data:", error);
        this.loading = false;
      }
    );
  }

  selectTopic(quizTopic: string): void {
    const encodedTopic = encodeURIComponent(quizTopic);
    this.quizTopic = quizTopic;
    this.loading = true;
    this.quizStarted = false;
    this.currentQuestionIndex = 0;
    this.selectedAnswer = null;
    this.score = 0;

    // Navigate to new quiz topic
    this.router.navigate(['/', this.category, this.subcategory, encodedTopic, 'takequiz']).then(() => {
      this.loadQuizData();
    });
  }

  startQuiz(): void {
    if (!this.quizData) {
      console.error("Quiz data is not loaded!");
      return;
    }
    this.quizStarted = true;
    this.currentQuestionIndex = 0;
    this.selectedAnswer = null;
    this.score = 0;

    this.remainingTime = 100; // reset time
    this.showTimeWarning = false;
    this.showWarning = false;
    this.startTimer();
  }

  selectAnswer(optionId: number): void {
    this.selectedAnswer = optionId;
  }

  nextQuestion(): void {
    if (this.quizData && this.currentQuestionIndex < this.quizData.questions.length - 1) {
      this.checkAnswer();
      this.currentQuestionIndex++;
      this.selectedAnswer = null;
    }
  }

  submitQuiz(): void {
    this.checkAnswer(); // Check final answer before scoring
    // Show score popup
    this.showModal = true;
  
    // Don’t reset quizTopic/quizData here — keep quiz on screen
    this.quizStarted = false;
  
    const userId = localStorage.getItem('userId'); 
    if (!userId) {
      console.error('User ID not found in local storage');
      return;
    }
  
    if (!this.quizData) {
      console.error('Quiz data is not available');
      return;
    }
  
    const payload = {
      user_id: userId,
      quiz_topic_id: this.quiz_topic_id,
      quiz_topic_name: this.quizData.quiz_topic,
      Score: this.score
    };
  
    this.takequizService.submitQuizResults(payload).subscribe({
      next: (response) => {
        console.log('Quiz results submitted successfully:', response);
      },
      error: (error) => {
        console.error('Error submitting quiz results:', error);
      }
    });
  }
  

  checkAnswer(): void {
    if (this.quizData && this.selectedAnswer === this.quizData.questions[this.currentQuestionIndex].correct_option) {
      this.score++;
    }
  }

  closeModal(): void {
    this.showModal = false;
    this.quizTopic = '';  
    this.quizData = null;  
    this.quizStarted = false;
    this.selectedAnswer = null;
    this.currentQuestionIndex = 0;
    this.score = 0;

    // Navigate back to the base URL of the subcategory
    this.router.navigate(['/takequiz', this.category, this.subcategory]);
  }

  retryQuiz(): void {
    this.currentQuestionIndex = 0;
    this.score = 0;
    this.selectedAnswer = null;
    this.showModal = false;
    this.quizStarted = true;
  }

  isLastQuestion(): boolean {
    return this.quizData ? this.currentQuestionIndex === this.quizData.questions.length - 1 : false;
  }

  progressWidth(): string {
    if (!this.quizData) return '0%';
    return `${((this.currentQuestionIndex + 1) / this.quizData.questions.length) * 100}%`;
  }

  getOptionLabel(index: number): string {
    return ['A', 'B', 'C', 'D'][index] || '';
  }

  // New method to start the timer
  showTimeWarning: boolean = false;  
startTimer() {
  const timer = setInterval(() => {
    if (this.remainingTime > 0) {
      this.remainingTime--;

      // Show warning when time reaches 1 minute
      if (this.remainingTime === 60) {
        this.showTimeWarning = true; 
      }

      
      this.showWarning = this.remainingTime <= 30; 
    } else {
      clearInterval(timer);  
      this.submitQuiz(); // Auto-submit the quiz
    }
  }, 1000); 
}

  

  // New method to format the remaining time
  getFormattedTime(): string {
    let minutes = Math.floor(this.remainingTime / 60);
    let seconds = this.remainingTime % 60;
    return `${minutes}:${seconds < 10 ? '0' + seconds : seconds}`;
  }

}