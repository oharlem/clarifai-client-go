Release v0.5.1 (2016-12-12) 
===
Minor cleanup release to updated docs and refactoring.


Release v0.5.0 (2016-12-11)
===

- Added model operations
- DEV: added serverReset() test helper, to re-init test server for identical routes.


Release v0.4.0 (2016-12-09)
===
* Major API refactoring / simplification 
* Added all input operations
* Examples should use 2 environment variables CLARIFAI_API_ID and CLARIFAI_API_SECRET for communications with your Clarifai API account.


Release v0.3.0 (2016-11-28)
===
* Major refactoring & cleanup.


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


Release v0.1.0 (2016-11-24) aka "Thanksgiving Waiting Time"
===
* Authentication
* Basic prediction functionality: request by model, URL and base64 images as inputs. 
