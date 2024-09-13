
# Technical Test ITONICS


Ce Projet est un test technique donné par ITONICS dont le but est de développer une API en Golang et d'y inclure une méthode POST pour céer une Brainee ainsi qu'une méthode GET pour obtenir une brainee précédemment ajoutée, en fonction de son ID.

Pour ce projet, j'ai décider d'utiliser le framework Fiber ainsi qu'une base de donnée MongoDB.

Je vous remercie sincèrement pour le temps que vous accorderez à l'évaluation de ce projet ainsi que pour l'intérêt que vous portez à mon profil.

Loïc Marlard


## Prérequis

- Golang version 1.23.1
- Postman ou Insomnia

Vous pouvez importer cette collection à postman ou à insomnia :
```
{"_type":"export","__export_format":4,"__export_date":"2024-09-13T00:04:54.836Z","__export_source":"insomnia.desktop.app:v10.0.0","resources":[{"_id":"req_ce9063d57f7f41f38b5c93aa34353f51","parentId":"wrk_98b03f85e5244499b4b301c9177a7658","modified":1726179757384,"created":1726143403786,"url":"http://localhost:5000/brainees","name":"Create Brainee","description":"","method":"POST","body":{"mimeType":"application/json","text":"{\n\t\"text\": \"Et si on pouvait se faire livrer à manger dans le train ? Prochaine arrêt : livraison :)\",\n\t\"author\": \"Lambert\",\n\t\"brand\": \"SNCF\"\n}"},"parameters":[],"headers":[{"name":"Content-Type","value":"application/json"},{"name":"User-Agent","value":"insomnia/10.0.0"}],"authentication":{},"metaSortKey":-1726175517267.5,"isPrivate":false,"pathParameters":[],"settingStoreCookies":true,"settingSendCookies":true,"settingDisableRenderRequestBody":false,"settingEncodeUrl":true,"settingRebuildPath":true,"settingFollowRedirects":"global","_type":"request"},{"_id":"wrk_98b03f85e5244499b4b301c9177a7658","parentId":null,"modified":1726143395802,"created":1726143395802,"name":"TechnicalTestITONICS","description":"","scope":"collection","_type":"workspace"},{"_id":"req_9be2e5c9bacf4bb6a64e7326244595b2","parentId":"wrk_98b03f85e5244499b4b301c9177a7658","modified":1726179760957,"created":1726173619337,"url":"http://localhost:5000/brainees/66e3557ae720138fb5532a83","name":"Update Brainee","description":"","method":"PUT","body":{"mimeType":"application/json","text":"{\n\t\"author\": \"Largon\"\n}"},"parameters":[],"headers":[{"name":"Content-Type","value":"application/json"},{"name":"User-Agent","value":"insomnia/10.0.0"}],"authentication":{},"metaSortKey":-1726175280026.1875,"isPrivate":false,"pathParameters":[],"settingStoreCookies":true,"settingSendCookies":true,"settingDisableRenderRequestBody":false,"settingEncodeUrl":true,"settingRebuildPath":true,"settingFollowRedirects":"global","_type":"request"},{"_id":"req_df12a0c4733c4b8885665fb7e5e3e7dd","parentId":"wrk_98b03f85e5244499b4b301c9177a7658","modified":1726179773231,"created":1726177415198,"url":"http://localhost:5000/brainees/66e34d03700ac94c6c3fcb9c","name":"Delete Brainee","description":"","method":"DELETE","body":{},"parameters":[],"headers":[{"name":"User-Agent","value":"insomnia/10.0.0"}],"authentication":{},"metaSortKey":-1726175161405.5312,"isPrivate":false,"pathParameters":[],"settingStoreCookies":true,"settingSendCookies":true,"settingDisableRenderRequestBody":false,"settingEncodeUrl":true,"settingRebuildPath":true,"settingFollowRedirects":"global","_type":"request"},{"_id":"req_4b57260e5aff430aa09d25fda9abf0b2","parentId":"wrk_98b03f85e5244499b4b301c9177a7658","modified":1726182957487,"created":1726182949181,"url":"http://localhost:5000/brainees","name":"Get Brainee by ID","description":"","method":"GET","body":{},"parameters":[],"headers":[{"name":"User-Agent","value":"insomnia/10.0.0"}],"authentication":{},"metaSortKey":-1726169114408.4688,"isPrivate":false,"pathParameters":[],"settingStoreCookies":true,"settingSendCookies":true,"settingDisableRenderRequestBody":false,"settingEncodeUrl":true,"settingRebuildPath":true,"settingFollowRedirects":"global","_type":"request"},{"_id":"req_360386f70cca4b13952ce15b38e058e9","parentId":"wrk_98b03f85e5244499b4b301c9177a7658","modified":1726182938260,"created":1726182938260,"url":"http://localhost:5000/brainees","name":"Get Brainees","description":"","method":"GET","body":{},"parameters":[],"headers":[{"name":"User-Agent","value":"insomnia/10.0.0"}],"authentication":{},"metaSortKey":-1726163067411.4062,"isPrivate":false,"pathParameters":[],"settingStoreCookies":true,"settingSendCookies":true,"settingDisableRenderRequestBody":false,"settingEncodeUrl":true,"settingRebuildPath":true,"settingFollowRedirects":"global","_type":"request"},{"_id":"req_48474a13c0c24696850e38d90278291f","parentId":"wrk_98b03f85e5244499b4b301c9177a7658","modified":1726183024428,"created":1726180868553,"url":"http://localhost:5000/brainees_author/Largon","name":"Get Brainees by Author","description":"","method":"GET","body":{},"parameters":[],"headers":[{"name":"User-Agent","value":"insomnia/10.0.0"}],"authentication":{},"metaSortKey":-1726151092037.9375,"isPrivate":false,"pathParameters":[],"settingStoreCookies":true,"settingSendCookies":true,"settingDisableRenderRequestBody":false,"settingEncodeUrl":true,"settingRebuildPath":true,"settingFollowRedirects":"global","_type":"request"},{"_id":"req_f7fca230dbfb44adb80990509add1f02","parentId":"wrk_98b03f85e5244499b4b301c9177a7658","modified":1726183034112,"created":1726182899681,"url":"http://localhost:5000/brainees_brand/SNCF","name":"Get Brainees by Brand","description":"","method":"GET","body":{},"parameters":[],"headers":[{"name":"User-Agent","value":"insomnia/10.0.0"}],"authentication":{},"metaSortKey":-1726139116664.4688,"isPrivate":false,"pathParameters":[],"settingStoreCookies":true,"settingSendCookies":true,"settingDisableRenderRequestBody":false,"settingEncodeUrl":true,"settingRebuildPath":true,"settingFollowRedirects":"global","_type":"request"},{"_id":"env_a4ffebd7ac586fb88aa3f65fef3b01d006607ed2","parentId":"wrk_98b03f85e5244499b4b301c9177a7658","modified":1726143395803,"created":1726143395803,"name":"Base Environment","data":{},"dataPropertyOrder":null,"color":null,"isPrivate":false,"metaSortKey":1726143395803,"_type":"environment"},{"_id":"jar_a4ffebd7ac586fb88aa3f65fef3b01d006607ed2","parentId":"wrk_98b03f85e5244499b4b301c9177a7658","modified":1726143395804,"created":1726143395804,"name":"Default Jar","cookies":[],"_type":"cookie_jar"}]}
```


## Installation

Pour installer le projet :

```bash
  go mod tidy
```

## Variables d'environnement

Pour lancer ce projet, vous aurez besoin de cette variable dans votre .env
(le repository étant privé j'ai push le .env pour plus de simplicité)

`MONGO_URI`=mongodb+srv://ITONICS:TTPSW@tt-itonics-db.yq1so.mongodb.net/?retryWrites=true&w=majority&appName=TT-ITONICS-DB

## Lancement

Pour lancer le projet :

```bash
  go run main.go
```


## API Reference

#### Créé une brainee

```http
  POST http://localhost:5000/brainees
```
| Parameter | Type     | Description                              |
|:----------| :------- |:-----------------------------------------|
| `text`    | `string` | **Requis**. Le texte de votre brainee    |
| `author`  | `string` | **Requis**. L'auteur de votre brainee  |
| `brand`   | `string` | **Requis**. La marque de votre brainee |

#### Éditer une brainee

```http
  PUT http://localhost:5000/brainees/:braineesId
```

| Parameter | Type     | Description                                |
|:----------| :------- |:-------------------------------------------|
| `text`    | `string` | **Facultatif**. Le texte de votre brainee  |
| `author`  | `string` | **Facultatif**. L'auteur de votre brainee  |
| `brand`   | `string` | **Facultatif**. La marque de votre brainee |

Vous pouvez éditer un ou plusieurs champs à la fois.

#### Supprimer une brainee

```http
  DELETE http://localhost:5000/brainees/:braineesId
```

#### Get une brainee avec son ID

```http
  GET http://localhost:5000/brainees/:braineesId
```

#### Get toutes les brainees d'un auteur

```http
  GET http://localhost:5000/brainees_author/:braineesAuthor
```

#### Get toutes les brainees d'une marque

```http
  GET http://localhost:5000/brainees_brand/:braineesBrand
```
#### Get toutes les brainees

```http
  GET http://localhost:5000/brainees
```