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
  imports: [CommonModule,RouterModule]
})
export class  TakeQuizComponent implements OnInit{
  // topics: Topic[] = [
  //   {
  //     name: 'Angular',
  //     questions: [
  //       { question: 'What is Angular?', options: ['Framework', 'Library', 'Plugin', 'Theme'], correctAnswer: 'Framework' },
  //       { question: 'Which language is used in Angular?', options: ['Java', 'Python', 'TypeScript', 'C#'], correctAnswer: 'TypeScript' },
  //       { question: 'What is CLI in Angular?', options: ['Command Line Interface', 'Component Library Interface', 'Common Logic Integration', 'None'], correctAnswer: 'Command Line Interface' },
  //       { question: 'What is used for Dependency Injection?', options: ['Services', 'Directives', 'Components', 'Modules'], correctAnswer: 'Services' },
  //       { question: 'Which decorator is used for a component?', options: ['@Injectable', '@Component', '@Directive', '@Pipe'], correctAnswer: '@Component' }
  //     ]
  //   },
  //   {
  //     name: 'TypeScript',
  //     questions: [
  //       { question: 'What is TypeScript?', options: ['Programming Language', 'Superset of JavaScript', 'Framework', 'None'], correctAnswer: 'Superset of JavaScript' },
  //       { question: 'Which keyword is used to define a variable in TypeScript?', options: ['var', 'let', 'const', 'All of these'], correctAnswer: 'All of these' },
  //       { question: 'Which company developed TypeScript?', options: ['Google', 'Microsoft', 'Facebook', 'Amazon'], correctAnswer: 'Microsoft' },
  //       { question: 'What is used for defining types?', options: ['Type Annotations', 'Interfaces', 'Both', 'None'], correctAnswer: 'Both' },
  //       { question: 'Which file extension is used for TypeScript?', options: ['.js', '.ts', '.tsx', '.json'], correctAnswer: '.ts' }
  //     ]
  //   },
  //   {
  //     name: 'JavaScript',
  //     questions: [
  //       { question: 'Which keyword is used to declare a constant in JavaScript?', options: ['var', 'let', 'const', 'static'], correctAnswer: 'const' },
  //       { question: 'Which symbol is used for comments in JavaScript?', options: ['//', '/* */', '#', '--'], correctAnswer: '//' },
  //       { question: 'What is used to handle asynchronous operations in JavaScript?', options: ['Promises', 'Callbacks', 'Async/Await', 'All of the above'], correctAnswer: 'All of the above' },
  //       { question: 'Which company developed JavaScript?', options: ['Google', 'Microsoft', 'Netscape', 'IBM'], correctAnswer: 'Netscape' },
  //       { question: 'Which method is used to convert a JSON string into an object?', options: ['JSON.parse()', 'JSON.stringify()', 'toObject()', 'parseJSON()'], correctAnswer: 'JSON.parse()' }
  //     ]
  //   },
  //   {
  //     name: 'React',
  //     questions: [
  //       { question: 'What is React?', options: ['Library', 'Framework', 'Language', 'Database'], correctAnswer: 'Library' },
  //       { question: 'What is JSX?', options: ['JavaScript XML', 'JavaScript Extension', 'JSON Syntax', 'None'], correctAnswer: 'JavaScript XML' },
  //       { question: 'Which hook is used for state management?', options: ['useState', 'useEffect', 'useContext', 'useMemo'], correctAnswer: 'useState' },
  //       { question: 'What is the virtual DOM?', options: ['A lightweight copy of the real DOM', 'A new programming language', 'A database', 'A styling tool'], correctAnswer: 'A lightweight copy of the real DOM' },
  //       { question: 'Which method is used to render elements in React?', options: ['render()', 'display()', 'show()', 'None'], correctAnswer: 'render()' }
  //     ]
  //   },
  //   {
  //     name: 'Node.js',
  //     questions: [
  //       { question: 'Which language is Node.js built on?', options: ['Python', 'JavaScript', 'Java', 'C++'], correctAnswer: 'JavaScript' },
  //       { question: 'What is the default package manager for Node.js?', options: ['npm', 'pip', 'yarn', 'composer'], correctAnswer: 'npm' },
  //       { question: 'Which module is used to handle HTTP requests in Node.js?', options: ['http', 'fs', 'path', 'url'], correctAnswer: 'http' },
  //       { question: 'What is an event loop in Node.js?', options: ['A process that waits for and handles events', 'A function call', 'A database query', 'None'], correctAnswer: 'A process that waits for and handles events' },
  //       { question: 'Which framework is widely used with Node.js?', options: ['Express.js', 'Django', 'Flask', 'Laravel'], correctAnswer: 'Express.js' }
  //     ]
  //   },
  //   {
  //     name: 'Python',
  //     questions: [
  //       { question: 'What is Python?', options: ['Programming Language', 'Library', 'Framework', 'Markup Language'], correctAnswer: 'Programming Language' },
  //       { question: 'Which keyword is used for functions in Python?', options: ['func', 'define', 'def', 'lambda'], correctAnswer: 'def' },
  //       { question: 'What data structure does a Python list use?', options: ['Array', 'Linked List', 'Hash Table', 'Tree'], correctAnswer: 'Array' },
  //       { question: 'Which library is used for data analysis?', options: ['Pandas', 'NumPy', 'Scikit-learn', 'All of the above'], correctAnswer: 'All of the above' },
  //       { question: 'Which operator is used for exponentiation in Python?', options: ['^', '**', 'exp()', '//'], correctAnswer: '**' }
  //     ]
  //   },
  //   {
  //     name: 'Django',
  //     questions: [
  //       { question: 'What is Django?', options: ['Framework', 'Library', 'Database', 'Language'], correctAnswer: 'Framework' },
  //       { question: 'Which language is used in Django?', options: ['Python', 'JavaScript', 'PHP', 'Java'], correctAnswer: 'Python' },
  //       { question: 'What is ORM in Django?', options: ['Object Relational Mapper', 'Object Request Manager', 'Object Real Model', 'None'], correctAnswer: 'Object Relational Mapper' },
  //       { question: 'Which command starts a Django project?', options: ['django-admin startproject', 'django start', 'create-django', 'new-django'], correctAnswer: 'django-admin startproject' },
  //       { question: 'What is the default database used in Django?', options: ['MySQL', 'PostgreSQL', 'SQLite', 'MongoDB'], correctAnswer: 'SQLite' }
  //     ]
  //   },
  //   {
  //     name: 'Machine Learning',
  //     questions: [
  //       { question: 'What is Machine Learning?', options: ['Data Science', 'AI technique', 'Programming Language', 'Database'], correctAnswer: 'AI technique' },
  //       { question: 'Which library is used for ML in Python?', options: ['Scikit-learn', 'TensorFlow', 'Keras', 'All of the above'], correctAnswer: 'All of the above' },
  //       { question: 'What is supervised learning?', options: ['Learning with labeled data', 'Learning without data', 'Learning with only input', 'None'], correctAnswer: 'Learning with labeled data' },
  //       { question: 'Which algorithm is used for classification?', options: ['Decision Tree', 'Linear Regression', 'K-Means', 'Apriori'], correctAnswer: 'Decision Tree' },
  //       { question: 'What is the loss function in ML?', options: ['A function to measure error', 'A data preprocessing method', 'A dataset split technique', 'None'], correctAnswer: 'A function to measure error' }
  //     ]
  //   }
  // ];

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


  constructor(private route: ActivatedRoute,private router: Router,private takequizService: TakequizService)
  {}
  
  ngOnInit(): void {
    // Retrieve route parameters
    this.route.params.subscribe(params => {
      this.category = params['category'];
      this.subcategory = params['subcategory'];
      this.quizTopic = params['quizTopic'];
      console.log(this.quizTopic);

      // Fetch quiz data from JSON file
      this.loadQuizData();
    });
  }

  loadQuizData() {
    this.takequizService.getQuizData().subscribe(data => {
      console.log("Loaded JSON Data:", data); 
      // Find the selected category
      const selectedCategory = data.find((c: any) => c.category === this.category);
      // if (!selectedCategory) return;
      if (!selectedCategory) {
        console.error("Category Not Found:", this.category);
        return;
      }

      // // Find the selected subcategory
      const selectedSubcategory = selectedCategory.sub_categories.find(
         (s: any) => s.sub_category === this.subcategory
       );
      // if (!selectedSubcategory) return;

      // Store all quiz topics under this subcategory (for left panel)
      this.quizTopicsList = selectedSubcategory.quiz_topics.map((q: any) => q.quiz_topic);

        // Find the selected quiz topic
        const selectedQuiz = selectedSubcategory.quiz_topics.find(
            (q: any) => q.quiz_topic === this.quizTopic
        );

        if (!selectedQuiz) {
            console.error("Quiz Topic Not Found:", this.quizTopic);
            return;
        }

        this.quizData = selectedQuiz; 
        this.loading = false;
    });
  }

  selectTopic(quizTopic: string): void {
    this.quizTopic = quizTopic;
    this.loading = true;
    this.quizStarted = false;
    this.currentQuestionIndex = 0;
    this.selectedAnswer = null;
    this.score = 0;

    // Navigate to new quiz topic
    this.router.navigate(['/', this.category, this.subcategory, this.quizTopic, 'takequiz']).then(() => {
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
    this.checkAnswer();
    this.showModal = true;
  }

  checkAnswer(): void {
    if (this.quizData && this.selectedAnswer === this.quizData.questions[this.currentQuestionIndex].correct_option) {
      this.score++;
    }
  }

  closeModal(): void {
    this.showModal = false;
    this.quizTopic = '';  // Reset selected quiz topic
    this.quizData = null;  // Clear quiz data
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
}
