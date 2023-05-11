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

