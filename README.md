# Nuku Hiva Monarch

This project is named after a bird. Anyway, this gets my saved posts on Reddit and displays it on a website interface.

## Webserver setup

- Start a project in a directory with ```yarn init```
- Add parcel with ```yarn add --dev parcel```
- Add vue with ```yarn add --dev vue```
- Create an ```src``` directory
- In ```src```
  - Create a ```components``` directory
  - Create an ```App.vue``` file
  - Create an ```index.ts``` file
  - Create an ```index.scss``` file
- Go to package.json
  - Remove the "main" field
  - Add a "source" field and set the source to index.html like this:

    ```
      "source": "./index.html",
    ```

  - Add a "scripts" {} field
    - Inside this field, add:

    ```
        "scripts": {
          "start": "parcel",
          "build": "parcel build"
        },
    ```

## Links

- [How to get more JSON Results](https://old.reddit.com/r/redditdev/comments/d7egb/how_to_get_more_json_results_i_get_only_30/)

## Planned Features

- Add Vue and front-end support
- Encryption of the link.txt file
  - Must require a user set password during runtime
- Database support for saved posts using MongoDB or that of the likes of SQL
