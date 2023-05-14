const { buildSchema } = require('graphql');

const schema = buildSchema(`
  type User {
    fullname: String
    username: String
    id: Int
  }

  type Timezone {
    name: String
  }

  type Alarm {
    id: Int
    title: String
    time: String
  }

  type Timer {
    id: Int
    time: String
  }

  type Event {
    title: String
    description: String
    start: String
    end: String
    allDay: Int
    location: String
  }

  type Query {
    timezones: [Timezone]
    alarms(user_id: Int!): [Alarm]
    events: [Event]
  }

  type Mutation {
    createUser(user_id: Int, username: String, fullname: String, hashed_password: String): String
    deleteUser(user_id: Int): String
    createAlarm(user_id: Int!, newTitle: String, newTime: String): String
    deleteAlarm(alarm_id: Int!): String
    updateAlarm(alarm_id: Int!, newTitle: String, newTime: String): String
  }
`);

module.exports = schema;

/*

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0100   178  100   178    0     0  52383      0 --:--:-- --:--:-- --:--:-- 59333
[{"id":1,"title":"Mi evento","description":"Descripción de mi evento","start":"2022-03-13T18:30:00.000Z","end":"2022-03-20T18:40:00.000Z","allDay":0,"location":"Mi ubicación"}]
*/
