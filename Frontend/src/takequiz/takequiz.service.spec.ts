import { TestBed } from '@angular/core/testing';
import { TakequizService } from './takequiz.service';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';

describe('TakequizService', () => {
  let service: TakequizService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [
        TakequizService,
        {
          provide: 'environment', // Mock the environment
          useValue: {
            protectedApiUrl: 'http://localhost:8080/api'
          }
        }
      ]
    });

    service = TestBed.inject(TakequizService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify(); // Verify no outstanding requests
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should fetch quiz data for a given topic', () => {
    const mockQuizData = { 
      quiz_topic: 'Graphs', 
      questions: [{ 
        question: 'What is a graph?', 
        options: ['A', 'B'], 
        correct_option: 1 
      }] 
    };

    service.getQuizData('Graphs').subscribe(data => {
      expect(data).toEqual(mockQuizData);
    });

    // Match the exact URL with the /api prefix
    const req = httpMock.expectOne('http://localhost:8080/api/quiz/Graphs');
    expect(req.request.method).toBe('GET');
    req.flush(mockQuizData);
  });

  it('should fetch quiz topics for a category and subcategory', () => {
    const mockTopics = { 
      quiz_topics: ['Topic1', 'Topic2'] 
    };

    service.getQuizTopics('Math', 'Algebra').subscribe(data => {
      expect(data).toEqual(mockTopics.quiz_topics);
    });

    const req = httpMock.expectOne(
      'http://localhost:8080/api/categories/Math/subcategories/Algebra/quiz_topics'
    );
    expect(req.request.method).toBe('GET');
    req.flush(mockTopics);
  });

  it('should submit quiz results', () => {
    const mockPayload = {
      user_id: '123',
      quiz_topic_id: '456',
      quiz_topic_name: 'Graphs',
      Score: 80
    };

    service.submitQuizResults(mockPayload).subscribe();

    const req = httpMock.expectOne('http://localhost:8080/api/submitquiz');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual(mockPayload);
    req.flush({});
  });

  it('should handle errors when fetching quiz data', () => {
    const consoleSpy = spyOn(console, 'error');
    
    service.getQuizData('InvalidTopic').subscribe({
      error: (err) => {
        expect(err).toBeTruthy();
      }
    });

    const req = httpMock.expectOne('http://localhost:8080/api/quiz/InvalidTopic');
    req.flush('Error', { 
      status: 404, 
      statusText: 'Not Found' 
    });

    expect(consoleSpy).toHaveBeenCalledWith('Error fetching quiz data:', jasmine.any(Object));
  });
});