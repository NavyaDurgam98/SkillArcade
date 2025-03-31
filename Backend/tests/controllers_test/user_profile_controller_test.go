package controllers_test

// import (
// 	"BACKEND/controllers"
// 	"BACKEND/services"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// // Setting up the Gin router
// func setupRouter() *gin.Engine {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	controllers.UserProfileRouter(router)
// 	return router
// }

// func TestGetUserProfile_Success(t *testing.T) {
// 	router := setupRouter()

// 	// Mocking the GetUserProfileService function
// 	services.GetUserProfileService = func(ctx context.Context, collection interface{}, userID string) (interface{}, error) {
// 		return map[string]interface{}{
// 			"username":  "john_doe",
// 			"firstname": "John",
// 			"lastname":  "Doe",
// 			"email":     "john.doe@example.com",
// 		}, nil
// 	}

// 	req, _ := http.NewRequest(http.MethodGet, "/userprofile?user_id=valid_user_id", nil)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	// Assertions
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.JSONEq(t, `{"username":"john_doe","firstname":"John","lastname":"Doe","email":"john.doe@example.com"}`, w.Body.String())
// }

// func TestGetUserProfile_MissingUserID(t *testing.T) {
// 	router := setupRouter()

// 	req, _ := http.NewRequest(http.MethodGet, "/userprofile", nil)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	// Assertions
// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// 	assert.JSONEq(t, `{"error":"User ID is required"}`, w.Body.String())
// }

// func TestGetUserProfile_InvalidUserIDFormat(t *testing.T) {
// 	router := setupRouter()

// 	// Mocking the GetUserProfileService function for invalid ID format
// 	services.GetUserProfileService = func(ctx context.Context, collection interface{}, userID string) (interface{}, error) {
// 		return nil, errors.New("invalid user ID format")
// 	}

// 	req, _ := http.NewRequest(http.MethodGet, "/userprofile?user_id=invalid_id", nil)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	// Assertions
// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// 	assert.JSONEq(t, `{"error":"Invalid user ID format"}`, w.Body.String())
// }

// func TestGetUserProfile_UserNotFound(t *testing.T) {
// 	router := setupRouter()

// 	// Mocking the GetUserProfileService function for user not found
// 	services.GetUserProfileService = func(ctx context.Context, collection interface{}, userID string) (interface{}, error) {
// 		return nil, errors.New("user not found")
// 	}

// 	req, _ := http.NewRequest(http.MethodGet, "/userprofile?user_id=non_existent_id", nil)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	// Assertions
// 	assert.Equal(t, http.StatusNotFound, w.Code)
// 	assert.JSONEq(t, `{"error":"User not found"}`, w.Body.String())
// }

// func TestGetUserProfile_InternalServerError(t *testing.T) {
// 	router := setupRouter()

// 	// Mocking the GetUserProfileService function for internal server error
// 	services.GetUserProfileService = func(ctx context.Context, collection interface{}, userID string) (interface{}, error) {
// 		return nil, errors.New("internal error")
// 	}

// 	req, _ := http.NewRequest(http.MethodGet, "/userprofile?user_id=some_user_id", nil)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	// Assertions
// 	assert.Equal(t, http.StatusInternalServerError, w.Code)
// 	assert.JSONEq(t, `{"error":"internal error"}`, w.Body.String())
// }
