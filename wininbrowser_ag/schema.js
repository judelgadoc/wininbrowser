const { buildSchema } = require('graphql');

const schema = buildSchema(`
  type Timezone {
    name: String
  }

  type Query {
    timezones: [Timezone]
  }
`);

module.exports = schema;

