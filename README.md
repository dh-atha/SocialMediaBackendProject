# SocialMediaBackendProject

## Base URL : https:/athaprojects.me

Working API:
POST /users (register)
POST /login
GET /users (getAllUsers)
GET /users/:id (getSpecificUser)
GET /profile (MyProfile)
PUT /profile
DELETE /profile
PUT /profilepic (updateProfilePic only)

GET /posts (getAllPosts)
GET /posts/:id (getSpecificPost with comments)
POST /myposts (addPost) (caption only/caption + post_images)
GET /myposts 
PUT /myposts/:id
DELETE /myposts/:id

POST /comments/:id (AddComment use param : post_id)
DELETE /comments/:id (DeleteComment use param : id (commentID))

Bugs:
Not found yet
