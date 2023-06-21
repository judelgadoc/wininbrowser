const express = require('express');
const cors = require( `cors` );
const https = require('https');
const fs = require('fs');
const {
    graphqlHTTP
} = require('express-graphql');
const fetch = require('node-fetch');
const schema = require('./schema');
var amqp = require('amqplib/callback_api');

const app = express();

const root = {
    createUser: async ({
        user_id,
        username,
        fullname,
        hashed_password
    }) => {
        const body = {
            id: parseInt(user_id),
            username: username,
            fullname: fullname,
            password: hashed_password
        }
        const response0 = await fetch(`http://wininbrowser-authentication-ms:8000/users`, {
            method: 'post',
            body: JSON.stringify(body),
            headers: {
                'Content-Type': 'application/json'
            }
        });
        const response1 = await fetch(`http://wininbrowser-clock-ms:9090/users/${user_id}`, {
            method: 'post',
        });
        const response2 = await fetch(`http://wininbrowser-calendar-ms:3000/users/${user_id}`, {
            method: 'post',
        });
        if ((response0.status === 200) && (response1.status === 200) && (response2.status === 200))
            return "User added successfully"
        else
            return "Error adding user"
    },
    deleteUser: async ({
        user_id
    }) => {
        const response0 = await fetch(`http://wininbrowser-authentication-ms:8000/users/${user_id}`, {
            method: 'delete'
        });
        const response1 = await fetch(`http://wininbrowser-clock-ms:9090/users/${user_id}`, {
            method: 'delete'
        });
        const response2 = await fetch(`http://wininbrowser-calendar-ms:3000/users/${user_id}`, {
            method: 'delete'
        });
        if ((response0.status === 200) && (response1.status === 200) && (response2.status === 200))
            return "User deleted successfully"
        else
            return "Error deleting user"
    },
    getToken: async ({username, password}) => {
        const response = await fetch(`http://wininbrowser-authentication-ms:8000/token`, {
            method: 'post',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
                'accept': 'application/json'
            },
            body: new URLSearchParams({
                'username': username,
                'password': password,
              })
        });
        const data = await response.json();
        return data
    },
    users: async () => {
        const response = await fetch(`http://wininbrowser-authentication-ms:8000/users`);
        const data = await response.json();
        if (data === null)
            return [{}]
        else
            return data.map((u) => ({
                username: u.username,
                fullname: u.fullname,
                id: u.id
            }));
    },
    userById: async ( {user_id} ) => {
        const response = await fetch(`http://wininbrowser-authentication-ms:8000/users/${user_id}`);
        const data = await response.json();
        return data
    },
    timezones: async () => {
        const response = await fetch('http://wininbrowser-clock-ms:9090/timezones');
        const data = await response.json();
        return data.map((timezone) => ({
            name: timezone.name
        }));
    },
    alarms: async ({
        user_id
    }) => {
        const response = await fetch(`http://wininbrowser-clock-ms:9090/${user_id}/alarms`);
        const data = await response.json();
        if (data === null)
            return [{}]
        else
            return data.map((alarm) => ({
                id: alarm.id,
                title: alarm.title,
                time: alarm.time
            }));
    },
    createAlarm: async ({
        user_id,
        newTitle,
        newTime
    }) => {
        const body = {
            title: newTitle,
            time: newTime
        }
        const response = await fetch(`http://wininbrowser-clock-ms:9090/${user_id}/alarms`, {
            method: 'post',
            body: JSON.stringify(body),
            headers: {
                'Content-Type': 'application/json'
            }
        });
        if (response.status === 200)
            return "Alarm added successfully"
        else
            return "Error adding alarm"
    },
    deleteAlarm: async ({
        alarm_id
    }) => {
        const response = await fetch(`http://wininbrowser-clock-ms:9090/alarms/${alarm_id}`, {
            method: 'delete'
        });
        if (response.status === 200)
            return "Alarm deleted successfully"
        else
            return "Error deleting alarm"
    },
    updateAlarm: async ({
        alarm_id,
        newTitle,
        newTime
    }) => {
        const body = {
            title: newTitle,
            time: newTime
        }
        const response = await fetch(`http://wininbrowser-clock-ms:9090/alarms/${alarm_id}`, {
            method: 'put',
            body: JSON.stringify(body),
            headers: {
                'Content-Type': 'application/json'
            }
        });
        if (response.status === 200)
            return "Alarm updated successfully"
        else
            return "Error updating alarm"
    },
    timers: async ({
        user_id
    }) => {
        const response = await fetch(`http://wininbrowser-clock-ms:9090/${user_id}/timers`);
        const data = await response.json();
        if (data === null)
            return [{}]
        else
            return data.map((timers) => ({
                id: timers.id,
                time: timers.time
            }));
    },
    events: async () => {
        let test;
        amqp.connect('amqp://wininbrowser-mq', function (error0, connection) {
            if (error0) {
                throw error0;
            }
            connection.createChannel(function (error1, channel) {
                if (error1) {
                    throw error1;
                }
                var queue = 'calendar_queue';

                channel.assertQueue(queue, {
                    durable: false
                });

                channel.consume(queue, function reply(msg) {

                    console.log("Mensaje recibido: ", msg.content.toString())
                    const data = JSON.stringify(msg.content.toString());
                    test = data;

                    channel.sendToQueue(msg.properties.replyTo,
                        Buffer.from("events"), {
                        correlationId: msg.properties.correlationId
                    });
                    console.log("Mensaje enviado: events")

                    channel.ack(msg);
                });
            });
        });
        const response = await fetch('http://wininbrowser-calendar-ms:3000/events');
        const data = await response.json();
        return data.map((event_) => ({
            id: event_.id,
            title: event_.title,
            description: event_.description,
            start: event_.start,
            end: event_.end,
            allDay: event_.allDay,
            location: event_.location
        }));
    },
    eventsByUserId: async ( {user_id} ) => {
        const response = await fetch(`http://wininbrowser-calendar-ms:3000/${user_id}/events`);
        const data = await response.json();
        return data.map((event_) => ({
            id: event_.id,
            title: event_.title,
            description: event_.description,
            start: event_.start,
            end: event_.end,
            allDay: event_.allDay,
            location: event_.location
        }));
    },
    eventsByUsername: async ( {username} ) => {
        const response0 = await fetch(`http://wininbrowser-authentication-ms:8000/usernames/${username}`);
        const data0 = await response0.json();
        console.log(data0)
        const user_id = data0.id;
        const response1 = await fetch(`http://wininbrowser-calendar-ms:3000/${user_id}/events`);
        const data1 = await response1.json();
        return data1.map((event_) => ({
            id: event_.id,
            title: event_.title,
            description: event_.description,
            start: event_.start,
            end: event_.end,
            allDay: event_.allDay,
            location: event_.location
        }));
    },
    createEvent: async ({ title, description, start, end, allDay, location, userId }) => {
        amqp.connect('amqp://wininbrowser-mq', function (error0, connection) {
            if (error0) {
                throw error0;
            }
            connection.createChannel(function (error1, channel) {
                if (error1) {
                    throw error1;
                }
                var queue = 'calendar_queue';

                channel.assertQueue(queue, {
                    durable: false
                });
                channel.consume(queue, function reply(msg) {

                    console.log(" [x] En efecto se recibió la respuesta Create", msg.content.toString())
                    

                    channel.sendToQueue(msg.properties.replyTo,
                        Buffer.from(JSON.stringify(body)), {
                        correlationId: msg.properties.correlationId
                    });
                    console.log(" [x]Sent: Mensaje recibido")

                    channel.ack(msg);
                });
            });
        });
        const body = { title: title, description: description, start: start, end: end, allDay: allDay, location: location, userId: userId }
        const response = await fetch('http://wininbrowser-calendar-ms:3000/events', {
            method: 'post',
            body: JSON.stringify(body),
            headers: {
                'Content-Type': 'application/json'
            }
        });
        if (response.status === 200)
            return "Event added successfully"
        else
            return "Error adding event"
    },
    deleteEvent: async ({
        event_id
    }) => {
        const response = await fetch(`http://wininbrowser-calendar-ms:3000/events/${event_id}`, {
            method: 'delete'
        });
        console.log(response)
        if ((response.status === 200) || (response.status === 204))
            return "Event deleted successfully"
        else
            return "Error deleting event"
    },
    updateEvent: async ({
        event_id, title, description, start, end, allDay, location
    }) => {
        const body = {
            title: title,
            description: description,
            start: start,
            end: end,
            allDay: allDay,
            location: location
        }
        const response = await fetch(`http://wininbrowser-calendar-ms:3000/events/${event_id}`, {
            method: 'put',
            body: JSON.stringify(body),
            headers: {
                'Content-Type': 'application/json'
            }
        });
        if (response.status === 200)
            return "Event updated successfully"
        else
            return "Error updating event"
    },
    interopWith1F: async () => {
        const response = await fetch(`http://wininbrowser-int:29162`);
        const data = await response.json();
        return data.map((scheduledPayment) => ({
            UserId: scheduledPayment.UserId,
            Name: scheduledPayment.Name,
            CategoryId: scheduledPayment.CategoryId,
            AccountId: scheduledPayment.AccountId,
            PaymentMethod: scheduledPayment.PaymentMethod,
            Recipient: scheduledPayment.Recipient,
            Frequency: scheduledPayment.Frequency,
            StartDate: scheduledPayment.StartDate,
            NotificationTime: scheduledPayment.NotificationTime
        }));
    },
    disks: async () => {
        const response = await fetch('http://172.17.0.3:8080/disk/all');
        const data = await response.json();
        return data.map((disk) => ({ name: disk.name, folders: disk.folders }));
    },

    foldersFromFolder: async ({ diskName, route }) => {
        var queryRoute = 'http://172.17.0.3:8080/folder/getFolders?diskName=' + diskName + '&route=' + route
        const response = await fetch(queryRoute);
        const data = await response.json();
        return data.map((folder) => ({ name: folder.name, folders: folder.folders }));
    },

    filesFromFolder: async ({ diskName, route }) => {
        var queryRoute = 'http://172.17.0.3:8080/folder/getFiles?diskName=' + diskName + '&route=' + route
        const response = await fetch(queryRoute);
        const data = await response.json();
        return data.map((folder) => ({ name: folder.name, folders: folder.folders }));
    },

    newDisk: async ({ name, maximumSize }) => {
        var queryRoute = 'http://172.17.0.3:8080/disk/new?name=' + name + '&maximumSize=' + maximumSize + '&size=' + 0
        const response = await fetch(queryRoute, {
            method: 'POST'
        });
        return "Done"
    },

    newFolderInDisk: async ({ diskName, name }) => {
        var queryRoute = 'http://172.17.0.3:8080/disk/newFolder?diskName=' + diskName + '&name=' + name
        const response = await fetch(queryRoute, {
            method: 'POST'
        });
        console.log(response)
        return "Done"
    },

    newFolderInFolder: async ({ route, name, diskName }) => {
        var queryRoute = 'http://172.17.0.3:8080/folder/newFolder?route=' + route + '&name=' + name + '&diskName=' + diskName
        const response = await fetch(queryRoute, {
            method: 'PUT'
        });
        console.log(response)
        return "Done"
    },

    newFile: async ({ name, diskName, route, type, size }) => {
        var queryRoute = 'http://172.17.0.3:8080/folder/newFile?name=' + name + '&diskName=' + diskName + '&route=' + route + '&type=' + type + '&size=' + size
        const response = await fetch(queryRoute, {
            method: 'PUT'
        });
        console.log(response)
        return "Done"
    },
};

app.use( cors() );
app.use(
    '/graphql',
    graphqlHTTP({
        schema,
        rootValue: root,
        graphiql: true,
    })
);


const options = {
    cert: fs.readFileSync('cert.pem'),
    key: fs.readFileSync('key.pem'),
    rejectUnauthorized: false
  };
  
  https.createServer(options, app).listen(4000, () => {
    console.log('Server running on port 4000');
  });

/*  
app.listen(4000, () => {
    console.log('Server running on port 4000');
});
*/
//[{"id":1,"title":"Mi evento","description":"Descripción de mi evento","start":"2022-03-13T18:30:00.000Z","end":"2022-03-20T18:40:00.000Z","allDay":0,"location":"Mi ubicación"}]
