import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { LeaderboardComponent } from './leaderboard.component'; // Import standalone component
import { LeaderboardService } from './leaderboard.service';
import { of, throwError } from 'rxjs';
import { By } from '@angular/platform-browser';

describe('LeaderboardComponent', () => {
  let component: LeaderboardComponent;
  let fixture: ComponentFixture<LeaderboardComponent>;
  let leaderboardService: LeaderboardService;

  const mockLeaderboardData = [
    { rank: 1, username: 'User1', total_score: 100, quizzes_taken: 10 },
    { rank: 2, username: 'User2', total_score: 90, quizzes_taken: 8 }
  ];

  const mockUserData = { rank: 1, quizzes_taken: 10 };

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, LeaderboardComponent],  // Add standalone component here
      providers: [LeaderboardService]
    }).compileComponents();
  
    fixture = TestBed.createComponent(LeaderboardComponent);
    component = fixture.componentInstance;
    leaderboardService = TestBed.inject(LeaderboardService);
  });

  it('should create the component', () => {
    expect(component).toBeTruthy();
  });

  it('should load leaderboard data on init', () => {
    spyOn(leaderboardService, 'getLeaderboard').and.returnValue(of(mockLeaderboardData));
    spyOn(leaderboardService, 'getUserRank').and.returnValue(of(mockUserData));
    
    component.ngOnInit();
    fixture.detectChanges();

    expect(component.leaderboardData).toEqual(mockLeaderboardData);
    expect(component.currentUserRank).toBe(1);
    expect(component.attemptsMade).toBe(10);
  });

  it('should handle API error gracefully', () => {
    spyOn(leaderboardService, 'getLeaderboard').and.returnValue(throwError(() => new Error('API error')));
    spyOn(leaderboardService, 'getUserRank').and.returnValue(of(mockUserData));
    
    component.ngOnInit();
    fixture.detectChanges();

    expect(component.leaderboardData).toEqual([]);
    expect(component.isLoading).toBeFalse();
  });

  it('should toggle sorting order on column click', () => {
    component.leaderboardData = [...mockLeaderboardData];
    component.toggleSort('total_score');
    expect(component.sortColumn).toBe('total_score');
    expect(component.sortDirection).toBeTrue();

    component.toggleSort('total_score');
    expect(component.sortDirection).toBeFalse();
  });

  it('should sort leaderboard data correctly', () => {
    component.leaderboardData = [...mockLeaderboardData];
    component.sortColumn = 'total_score';
    component.sortDirection = true;
    component.sortLeaderboard();
    expect(component.leaderboardData[0].total_score).toBe(90);
  });

  it('should display leaderboard data in the table', () => {
    // Set up spy to return mock data
    spyOn(leaderboardService, 'getLeaderboard').and.returnValue(of(mockLeaderboardData));
    spyOn(leaderboardService, 'getUserRank').and.returnValue(of(mockUserData));

    // Trigger ngOnInit to load the data
    component.ngOnInit();
    
    // Trigger change detection to update the view
    fixture.detectChanges();

    // Query for all rows in the table
    const rows = fixture.debugElement.queryAll(By.css('tbody tr'));
    
    // Verify the correct number of rows
    expect(rows.length).toBe(mockLeaderboardData.length);

    // Optionally, check the contents of the first row
    const firstRowColumns = rows[0].queryAll(By.css('td'));
    expect(firstRowColumns[0].nativeElement.textContent).toBe('1'); // Rank
    expect(firstRowColumns[1].nativeElement.textContent).toBe('User1'); // Username
    expect(firstRowColumns[2].nativeElement.textContent).toBe('100'); // Total Score
    expect(firstRowColumns[3].nativeElement.textContent).toBe('10'); // Quizzes Taken
  });
});
