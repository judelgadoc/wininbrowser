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
  alarms: async ({ user_id }) => {
    const response = await fetch(`http://172.17.0.1:9090/${user_id}/alarms`);
    const data = await response.json();
      if (data === null) 
        return [{}]
      else
        return data.map((alarm) => ({ id: alarm.id, title: alarm.title, time: alarm.time }));
  },
  createAlarm: async ({ user_id, newTitle, newTime }) => {
    const body = {title: newTitle, time: newTime}
      console.log(user_id)
      console.log(body)
      console.log(JSON.stringify(body))
    const response = await fetch(`http://172.17.0.1:9090/${user_id}/alarms`, {
        method: 'post',
        body: JSON.stringify(body),
	    headers: {'Content-Type': 'application/json'}
    });
      if (response.status === 200) 
        return "Alarm added successfully"
      else
        return "Error adding alarm"
  },
  deleteAlarm: async ({ alarm_id }) => {
    const response = await fetch(`http://172.17.0.1:9090/alarms/${alarm_id}`, {
        method: 'delete'
    });
      if (response.status === 200) 
        return "Alarm deleted successfully"
      else
        return "Error deleting alarm"
  },
  updateAlarm: async ({ alarm_id, newTitle, newTime }) => {
    const body = {title: newTitle, time: newTime}
    const response = await fetch(`http://172.17.0.1:9090/alarms/${alarm_id}`, {
        method: 'put',
        body: JSON.stringify(body),
	    headers: {'Content-Type': 'application/json'}
    });
      if (response.status === 200) 
        return "Alarm updated successfully"
      else
        return "Error updating alarm"
  },
  timers: async ({ user_id }) => {
    const response = await fetch(`http://172.17.0.1:9090/${user_id}/timers`);
    const data = await response.json();
      if (data === null) 
        return [{}]
      else
        return data.map((timers) => ({ id: timer.id, time: timer.time }));
  },
  events: async () => {
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
