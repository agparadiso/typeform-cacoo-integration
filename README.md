# typeform-cacoo-integration
Typeform integration for Cacoo

This integration has been build in order to get quick and effective feedback on your cacoo diagrams by using the power of Typeform.

The way it works is: it gets cacoo diagram image, uploads it to typeform images, creates a form with the questions entered adding the diagram image in every question.

In order to run this integration in local just

`go run cmd/tfcacoo/main.go`

Now you will have the integration listening on port 3000 waiting for your call to build the typeform

So before being able to get the forms you need a Typeform account and to get your AccessToken (more info: https://developer.typeform.com/get-started/personal-access-token/)
I already assume you also have a cacoo account and a diagramID you would love to get feedback, (for this integration we use the cacooApiKey)

So to query the integration using httpie you would enter your questions separated by `,`:

`http localhost:3000/api/v1/getfeedback questions=="Do you think this diagram would be interpretated by a new team member?, What would you change?" tfapikey=="yourTypeformAccessToken" cacooapikey=="yourCacooApiKey" diagramID=="aCacooDiagramID"`

you can use this typeform test account token: `3EMKgVufspT9iqEsXxhTsWS3KbY754bm7VMTp9DkVtoa`

You should get as a response a link to the new typeform with your diagram, something like this:

{
    "Tflink": "https://tfintegration.typeform.com/to/SomeFormID"
}

you can now share this link to get feedback on your diagram, to check the responses log in into your tf account and enjoy!

tf test account:

email: xefu@utooemail.com
pass: tfintegration

Things I would improve with more time:
- Error handling
- Test coverage
- Add more question blocks
- Build deploy pipeline and deploy it
