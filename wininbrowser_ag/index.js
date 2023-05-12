const express = require('express');
const { graphqlHTTP } = require('express-graphql');
const fetch = require('node-fetch');
const schema = require('./schema');

const app = express();

const root = {
  timezones: async () => {
    const response = await fetch('http://172.17.0.1:9090/timezones');
    const data = await response.json();
    return data.map((timezone) => ({ name: timezone.name }));
  },
  aguacates: async () => {
    const response = await fetch('http://172.17.0.1:3000/events');
    const data = await response.json();
    return data.map((event_) => ({ title: event_.title, description: event_, start: event_.start, end: event_.end, allDay: event_.allDay, location: event_.location}));
  },
};

app.use(
  '/graphql',
  graphqlHTTP({
    schema,
    rootValue: root,
    graphiql: true,
  })
);

app.listen(4000, () => {
  console.log('Server running on port 4000');
});

//[{"id":1,"title":"Mi evento","description":"Descripción de mi evento","start":"2022-03-13T18:30:00.000Z","end":"2022-03-20T18:40:00.000Z","allDay":0,"location":"Mi ubicación"}]
