import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { QuizService } from './quiz.service';

// TypeScript Interface
interface QuizTopic {
  quiz_topic_id: string;
  quiz_topic_name: string;
  quizImgPath: string;
}

@Component({
  selector: 'app-quiz',
  standalone: true,
  templateUrl: './quiz.component.html',
  styleUrls: ['./quiz.component.css'],
  imports: [CommonModule]
})
export class QuizComponent implements OnInit {
  categoryName: string = '';
  subCategoryName: string = '';
  quizTopics: QuizTopic[] = [];

  constructor(private route: ActivatedRoute, private router: Router, private quizService: QuizService) {}

  ngOnInit() {
    this.categoryName = this.route.snapshot.paramMap.get('category') || '';
    this.subCategoryName = this.route.snapshot.paramMap.get('sub_category') || '';

    this.quizService.getQuizTopics(this.categoryName, this.subCategoryName).subscribe({
      next: (response: any) => {
        if (response && response.quiz_topics) {
          this.quizTopics = response.quiz_topics;
          console.log(`Quiz topics for ${this.subCategoryName}:`, this.quizTopics);
        } else {
          console.warn(`No quiz topics found for ${this.subCategoryName}`);
        }
      },
      error: (error) => {
        console.error('Error fetching quiz topics:', error);
      },
      complete: () => {
        console.log('Quiz topics fetch completed');
      }
    });
  }

  takeQuiz(quizTopicId: string) {
    console.log(`Navigating to: /${this.categoryName}/${this.subCategoryName}/${quizTopicId}/takequiz`);
    this.router.navigate([`/${this.categoryName}/${this.subCategoryName}/${quizTopicId}/takequiz`]);
  }
}
