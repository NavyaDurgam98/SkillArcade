package services

import (
	"BACKEND/Data"
	"context"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SearchQuizzes searches MongoDB for quiz topics only
func SearchQuizzes(searchText, page string) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Data.GetCollection("SkillArcade", "Quizzes")

	searchPattern := primitive.Regex{Pattern: ".*" + strings.ToLower(searchText) + ".*", Options: "i"}

	var filter bson.M
	if page == "quiz_topics" {
		filter = bson.M{"sub_categories.quiz_topics.quiz_topic_name": searchPattern}
	} else {
		return nil, nil
	}

	cursor, err := collection.Find(ctx, filter, options.Find().SetLimit(10))
	if err != nil {
		log.Println("Error searching quizzes:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	//  Extract all quiz_topics in the same sub-category as the matched one
	updatedResults := []bson.M{}
	for _, doc := range results {
		subCategories, ok := doc["sub_categories"].(primitive.A) //  Use `primitive.A` for BSON arrays
		if !ok {
			continue
		}

		//  Iterate over each sub-category
		for _, subCat := range subCategories {
			subCategory, ok := subCat.(bson.M)
			if !ok {
				continue
			}

			quizTopics, ok := subCategory["quiz_topics"].(primitive.A) //  Fix type assertion for array
			if !ok {
				continue
			}

			//  Find matching quiz topic and return all topics in the same sub-category
			found := false
			for _, quiz := range quizTopics {
				quizTopic, ok := quiz.(bson.M)
				if !ok {
					continue
				}

				if strings.EqualFold(quizTopic["quiz_topic_name"].(string), searchText) {
					found = true
					break
				}
			}

			//  If a match is found, add all quiz topics in the sub-category
			if found {
				for _, quiz := range quizTopics {
					quizTopic, ok := quiz.(bson.M)
					if !ok {
						continue
					}

					updatedResults = append(updatedResults, bson.M{
						"quiz_topic_id":   quizTopic["quiz_topic_id"],
						"quiz_topic_name": quizTopic["quiz_topic_name"],
					})
				}
			}
		}
	}

	return updatedResults, nil
}

// package services

// import (
// 	"BACKEND/Data"
// 	"context"
// 	"log"
// 	"strings"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// // SearchQuizzes searches MongoDB for categories, subcategories, or quizzes
// func SearchQuizzes(searchText, page string) ([]bson.M, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	collection := Data.GetCollection("SkillArcade", "Quizzes") //  Ensure correct collection

// 	searchPattern := primitive.Regex{Pattern: ".*" + strings.ToLower(searchText) + ".*", Options: "i"}

// 	var filter bson.M

// 	if page == "category" {
// 		filter = bson.M{"category": searchPattern}
// 	} else if page == "subcategory" {
// 		filter = bson.M{"sub_categories.sub_category": searchPattern}
// 	} else if page == "quiz" {
// 		//  Fix: Ensure correct query for searching nested quiz topics
// 		filter = bson.M{"sub_categories.quiz_topics.quiz_topic_name": searchPattern}
// 	} else {
// 		return nil, nil
// 	}

// 	cursor, err := collection.Find(ctx, filter, options.Find().SetLimit(100))
// 	if err != nil {
// 		log.Println("Error searching quizzes:", err)
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	var results []bson.M
// 	if err = cursor.All(ctx, &results); err != nil {
// 		return nil, err
// 	}

// 	//  Fix: If searching for a quiz topic, return full sub-category's quiz topics
// 	if page == "quiz" {
// 		updatedResults := []bson.M{}
// 		for _, doc := range results {
// 			subCategories, ok := doc["sub_categories"].([]interface{})
// 			if !ok {
// 				continue
// 			}

// 			for _, subCat := range subCategories {
// 				subCategory, ok := subCat.(bson.M)
// 				if !ok {
// 					continue
// 				}

// 				quizTopics, ok := subCategory["quiz_topics"].([]interface{})
// 				if !ok {
// 					continue
// 				}

// 				var matchedQuiz bson.M
// 				for _, quiz := range quizTopics {
// 					quizTopic, ok := quiz.(bson.M)
// 					if !ok {
// 						continue
// 					}

// 					//  Ensure search is case-insensitive
// 					if strings.Contains(strings.ToLower(quizTopic["quiz_topic_name"].(string)), strings.ToLower(searchText)) {
// 						matchedQuiz = quizTopic
// 					}
// 				}

// 				//  Append result if a match is found
// 				if matchedQuiz != nil {
// 					updatedResults = append(updatedResults, bson.M{
// 						"category":     doc["category"],
// 						"sub_category": subCategory["sub_category"],
// 						"quiz_topics":  quizTopics, //  Return full quiz topic list in the sub-category
// 						"matched_quiz": matchedQuiz,
// 					})
// 				}
// 			}
// 		}
// 		return updatedResults, nil
// 	}

// 	return results, nil
// }

// // GET http://localhost:8080/search?searchText=Com&page=category -- working

// // GET http://localhost:8080/search?searchText=Data&page=subcategory

// // GET http://localhost:8080/search?searchText=Java&page=quiz
