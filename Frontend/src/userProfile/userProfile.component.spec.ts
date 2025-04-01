import { ComponentFixture, TestBed } from '@angular/core/testing';
import { UserProfileComponent } from './userProfile.component';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { UserProfile, LeaderboardStats } from './userProfile.component';

describe('UserProfileComponent', () => {
  let component: UserProfileComponent;
  let fixture: ComponentFixture<UserProfileComponent>;
  let httpMock: HttpTestingController;

  const mockProfile: UserProfile = {
    username: 'testuser',
    firstname: 'Test',
    lastname: 'User',
    email: 'test@example.com',
    quizzes_taken: 5,
    rank: 10
  };

  const mockStats: LeaderboardStats = {
    username: 'testuser',
    total_score: 500,
    quizzes_taken: 5,
    rank: 10
  };

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        UserProfileComponent,
        HttpClientTestingModule,
        RouterTestingModule
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(UserProfileComponent);
    component = fixture.componentInstance;
    httpMock = TestBed.inject(HttpTestingController);

    // Mock localStorage
    spyOn(localStorage, 'getItem').and.returnValue('123');
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should initialize with default profile', () => {
    component.initializeProfile();
    expect(component.userProfile).toEqual({
      username: ' Not Found',
      firstname: ' Not Found',
      lastname: ' Not Found',
      email: ' Not Found',
      quizzes_taken: ' Not Found',
      rank: ' Not Found'
    });
  });

  it('should load user profile successfully', () => {
    component.loadUserProfile();

    const profileReq = httpMock.expectOne('http://localhost:8080/userprofile?user_id=123');
    expect(profileReq.request.method).toBe('GET');
    profileReq.flush(mockProfile);

    const statsReq = httpMock.expectOne('http://localhost:8080/leaderboard?user_id=123');
    expect(statsReq.request.method).toBe('GET');
    statsReq.flush(mockStats);

    expect(component.isLoading).toBeFalse();
    expect(component.userProfile).toEqual({
      ...mockProfile,
      quizzes_taken: mockStats.quizzes_taken,
      rank: mockStats.rank
    });
  });

  it('should handle profile load error', () => {
    component.loadUserProfile();

    const profileReq = httpMock.expectOne('http://localhost:8080/userprofile?user_id=123');
    profileReq.error(new ErrorEvent('Network error'));

    expect(component.isLoading).toBeFalse();
    expect(component.profileLoadError).toBeTrue();
  });

  it('should handle stats load error', () => {
    component.loadUserProfile();

    const profileReq = httpMock.expectOne('http://localhost:8080/userprofile?user_id=123');
    profileReq.flush(mockProfile);

    const statsReq = httpMock.expectOne('http://localhost:8080/leaderboard?user_id=123');
    statsReq.error(new ErrorEvent('Network error'));

    expect(component.isLoading).toBeFalse();
    expect(component.statsLoadError).toBeTrue();
  });

  it('should handle missing user ID', () => {
    (localStorage.getItem as jasmine.Spy).and.returnValue(null);
    component.loadUserProfile();

    expect(component.isLoading).toBeFalse();
    expect(component.profileLoadError).toBeTrue();
    httpMock.expectNone('http://localhost:8080/userprofile?user_id=');
  });

  it('should use default values when stats are missing', () => {
    const incompleteStats = {
      username: 'testuser',
      total_score: 500
    } as LeaderboardStats;

    component.loadUserProfile();

    const profileReq = httpMock.expectOne('http://localhost:8080/userprofile?user_id=123');
    profileReq.flush(mockProfile);

    const statsReq = httpMock.expectOne('http://localhost:8080/leaderboard?user_id=123');
    statsReq.flush(incompleteStats);

    expect(component.userProfile?.quizzes_taken).toBe('Not Found');
    expect(component.userProfile?.rank).toBe('Not Found');
  });
});