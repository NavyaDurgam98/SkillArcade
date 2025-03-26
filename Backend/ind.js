// const mongoose = require('mongoose');

// const uri = "mongodb+srv://baddamtejasri:skillarcade@cluster0.cw3xm.mongodb.net/SkillArcade";  // Ensure database name is included

// async function updateDocuments() {
//   try {
//     await mongoose.connect(uri);

//     const Category = mongoose.model('Category', new mongoose.Schema({}, { strict: false }), "Quizzes"); // Explicitly set "test" as the collection

//     // const result = await Category.updateMany({}, [
//     //   {
//     //     $set: {
//     //       "imgPath": "img1.png",
//     //       "sub_categories": {
//     //         $map: {
//     //           input: "$sub_categories",
//     //           as: "subCat",
//     //           in: {
//     //             $mergeObjects: [
//     //               "$$subCat",
//     //               { "subImgPath": "img3.png" },
//     //               {
//     //                 "quiz_topics": {
//     //                   $map: {
//     //                     input: "$$subCat.quiz_topics",
//     //                     as: "quiz",
//     //                     in: {
//     //                       $mergeObjects: [
//     //                         "$$quiz",
//     //                         { "quizImgPath": "img3.png" }
//     //                       ]
//     //                     }
//     //                   }
//     //                 }
//     //               }
//     //             ]
//     //           }
//     //         }
//     //       }
//     //     }
//     //   }
//     // ]);
//     const result = await Category.updateMany({}, [
//         {
//           $set: {
//             "sub_categories": {
//               $map: {
//                 input: "$sub_categories",
//                 as: "subCat",
//                 in: {
//                   $mergeObjects: [
//                     "$$subCat",
//                     { "subImgPath": "img2.png" }
//                   ]
//                 }
//               }
//             }
//           }
//         }
//       ]);

//     console.log(`${result.modifiedCount} documents updated successfully!`);
//   } catch (error) {
//     console.error("Error updating documents:", error);
//   } finally {
//     mongoose.disconnect();
//   }
// }

// updateDocuments();
const mongoose = require('mongoose');

// Replace with your actual MongoDB connection string
const mongoURI = "mongodb+srv://baddamtejasri:skillarcade@cluster0.cw3xm.mongodb.net/SkillArcade";


// Connect to MongoDB
mongoose.connect(mongoURI, { useNewUrlParser: true, useUnifiedTopology: true })
    .then(() => console.log("Connected to MongoDB!"))
    .catch(err => console.error("MongoDB Connection Error:", err));

// Define Mongoose Schema (Flexible)
const QuizSchema = new mongoose.Schema({}, { strict: false });
const Quiz = mongoose.model("Quizzes", QuizSchema, "Quizzes"); // Use correct collection name

// Search for "Java" in quiz_topic_name
async function searchQuizTopics() {
    try {
        const results = await Quiz.find({
            "sub_categories.quiz_topics.quiz_topic_name": { $regex: "Java", $options: "i" }
        }).limit(10);

        console.log("Search Results:", JSON.stringify(results, null, 2));
    } catch (err) {
        console.error("Error searching:", err);
    } finally {
        mongoose.disconnect();
    }
}

// Run the search
searchQuizTopics();
