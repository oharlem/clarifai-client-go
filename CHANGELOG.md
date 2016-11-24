Release v0.1.0 (2016-11-24) aka "Thanksgiving Waiting Time"
===
* Authentication
* Basic prediction functionality: request by model, URL and base64 images as inputs. 

Release v0.2.0 (2016-11-28)
===
* Added automatic token refresh on expiry.
* Input: 
    - get list of all inputs
* Search: 
    - Add images to a search index
    - Search by predicted concepts
    - Search by user supplied concept
    - Reverse image search
    - Search by custom metadata
    - Mixed search by concepts and predictions
* Pagination
* Image type validation as per [https://developer-preview.clarifai.com/guide/#supported-types](https://developer-preview.clarifai.com/guide/#supported-types)
* Refactoring
* Added examples into examples/ directory.