const fs = require('fs');
var faker = require('faker');
var DateGenerator = require('random-date-generator');

//Random Date starting date and ending date
let startDate = new Date(2019, 1, 2);
let endDate = new Date(2019, 3, 31);


//Random number generator
function getRandomInt(min, max) {
  min = Math.ceil(min);
  max = Math.floor(max);
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

//Variables set up to call each random value required in JSON data export
var randomDate = new Date(DateGenerator.getRandomDateInRange(startDate, endDate));
var randomName = faker.name.firstName() + " " + faker.name.lastName();
var randomBody = faker.lorem.paragraph();
var randomWords = faker.lorem.words(8);
var randomId = getRandomInt(150110000, 150119999);

//Piecing together the values and formatting to get ready for the squeeze
var randomTicket = {
  "Title": randomName + "-" + randomId + "-" + randomWords,
  "Created Time": "2019-01-01T22:44:32.237Z",
  "Updated Time": randomDate,
  "Body": randomBody
};

//Juicing the data into a tall glass of JSON
let data = JSON.stringify(randomTicket);
fs.writeFileSync('PROJ-001.json', data);

console.log("JSON created. Thats the juice.");

