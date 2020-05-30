# Playground for Swagger Editor

```shell-session
docker pull swaggerapi/swagger-editor
docker run -d -p 80:8080 swaggerapi/swagger-editor
```

Access to <http://localhost/>


## Notes

Track tutorial by seeing [here](https://medium.com/@amirm.lavasani/restful-apis-tutorial-of-openapi-specification-eeada0e3901d).

Master description is [OpenAPI Object](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md#schema).

- Metadata  Section

`openapi: 3.0.2`

- The Info Object

The information consists of
title, description, termsOfService, contact, license and version information.

<details>
  <summary>example</summary>

  ```yaml
  info:
    title: Blog Posts API
    description: >
      This is an example API for blog posts using OpenApi Specification.
      
      ---
      
      Using this API you can retrieve blog posts, comments on each blog
      post and delete or update the posts.
    termsOfService: "http://swagger.io/terms/"
    contact:
      name: Amir Lavasani
      url: "https://amir.lavasani.dev"
      email: amirm.lavasani@gmail.com
    license:
      name: "Apache 2.0"
      url: "http://www.apache.org/licenses/LICENSE-2.0.html"
    version: 1.0.0
  ```
</details>


- The Servers Object

The server consists of url and variables information.

<details>
  <summary>example</summary>
  
  ```yaml
  servers:
    - url: '{protocol}://{environment}.example.com/v1'
      variables:
        environment:
          default: api    # Production server
          enum:
            - api         # Production server
            - api.dev     # Development server
            - api.staging # Staging server
        protocol:
          default: https
          enum:
            - http
            - https
  ```
</details>

- The Paths Object

The path object conists of url key and method information value.
The method infomation can manage responses list according to each resopnse codes.

<details>
  <summary>example</summary>

  ```yaml
  paths:
    /posts:
  
      get:
        tags:
          - Posts
        summary: Return all the posts
        description: Return all the posts that are in our blog.
        responses:
          '200':
            description: An array of blog posts in JSON format
            content:
              application/json:
                schema: 
                  $ref: '#/components/schemas/Posts'
  
      post:
  ```
</details>

- The Components Object

<details>
  <summary>example</summary>
  
  ```yaml
  components:
    schemas:
      Post:
        type: object
        properties:
          id:
            type: string
          userId:
            type: string
          title:
            type: string
          body:
            type: string
        required:
          - id
          - userId
          - title
          - body
      Posts:
        type: array
        items:
          $ref: '#/components/schemas/Post'
  ```
</details>
