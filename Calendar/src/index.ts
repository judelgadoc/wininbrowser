import mysql, {
    RowDataPacket
} from 'mysql2';
import express, {
    Request,
    Response
} from 'express';
import * as amqp from 'amqplib/callback_api';

// Configuración de la conexión a la base de datos MySQL
const db = mysql.createPool({
    connectionLimit: 10,
    host: 'caldb',
    port: 65001,
    user: 'fredy',
    password: '12345',
    database: 'wininbrowser_calendar_db',
    debug: false
});

db.getConnection((err) => {
    if (err) {
        console.error('Error de conexión: ', err);
        return;
    }
    console.log('Conectado a la base de datos MySQL');
});

// Crear una instancia de Express
const app = express();

// Middleware para permitir el uso de JSON en el cuerpo de la solicitud
app.use(express.json());

function generateUuid() {
    return Math.random().toString() +
        Math.random().toString() +
        Math.random().toString();
}

function mqGetEventsClient() {
    let str : string;
    amqp.connect('amqp://wininbrowser-mq', function(error0, connection) {
        if (error0) {
            throw error0;
        }
        connection.createChannel(function(error1, channel) {
            if (error1) {
                throw error1;
            }

            var queue = 'calendar_queue';

            channel.assertQueue('', {
                exclusive: true
            },  async function(error2, q) {
                if (error2) {
                    throw error2;
                }
                var correlationId = generateUuid();

                channel.consume(q.queue, function(msg) {
                    if (msg !== null && msg !== undefined) {
                        if (msg.properties.correlationId == correlationId) {
                            console.log(" [x] Received %s", msg.content.toString());
                            str = msg.content.toString();

                        }
                    }
                }, {
                    noAck: true
                });
                const myans = await getAllEvents();
                
                channel.sendToQueue('calendar_queue',
                Buffer.from(JSON.stringify(myans)), {
                    correlationId: correlationId,
                    replyTo: q.queue
                });
                console.log(" [x] Sent: ", JSON.stringify(myans))
            });
        });
    });
}

function mqCreateEventClient() {
    amqp.connect('amqp://wininbrowser-mq', function(error0, connection) {
        if (error0) {
            throw error0;
        }
        connection.createChannel(function(error1, channel) {
            if (error1) {
                throw error1;
            }

            var queue = 'calendar_queue';

            channel.assertQueue('', {
                exclusive: true
            },  async function(error2, q) {
                if (error2) {
                    throw error2;
                }
                var correlationId = generateUuid();

                channel.consume(q.queue, async function(msg) {
                    if (msg !== null && msg !== undefined) {
                        if (msg.properties.correlationId == correlationId) {
                            console.log(" [x] Received %s", msg.content.toString());
                            const data = JSON.parse(msg.content.toString());
                            //const ans = await createEvent(data.title, data.description, data.start, data.end, data.allDay, data.location, data.userId);
                        }
                    }
                }, {
                    noAck: true
                });
                channel.sendToQueue('calendar_queue',
                Buffer.from("Event created correctly"), {
                    correlationId: correlationId,
                    replyTo: q.queue
                });
                console.log(" [x] Sent create");
            });
        });
    });
}

// Ruta para obtener todos los eventos
app.get('/events', async (req: Request, res: Response) => {
    try {
        const events = await getAllEvents();
        res.json(events);
        mqGetEventsClient()
    } catch (error) {
        res.status(500).send(error);
    }
});

// Ruta para obtener todos los eventos de un usuario
app.get('/:userId/events/', async (req: Request, res: Response) => {
    try {
        const {
            userId
        } = req.params;
        const events = await getAllEventsByUserId(parseInt(userId));
        res.json(events);
    } catch (error) {
        res.status(500).send(error);
    }
});


// Ruta para crear un nuevo evento
app.post('/events', async (req: Request, res: Response) => {
    try {
        const {
            title,
            description,
            start,
            end,
            allDay,
            location,
            userId
        } = req.body;
        const result = await createEvent(title, description, start, end, allDay, location, userId);
        res.json({
            title,
            start,
            end,
            allDay,
            location
        });
    } catch (error) {
        res.status(500).send(error);
    }
});

// Ruta para actualizar un evento existente
app.put('/events/:id', async (req: Request, res: Response) => {
    try {
        const {
            id
        } = req.params;
        const {
            title,
            description,
            start,
            end,
            allDay,
            location
        } = req.body;
        await updateEvent(parseInt(id), title, description, start, end, allDay, location);
        res.json({
            id: parseInt(id),
            title,
            start,
            end,
            allDay,
            location
        });
    } catch (error) {
        res.status(500).send(error);
    }
});

// Ruta para eliminar un evento existente
app.delete('/events/:id', async (req: Request, res: Response) => {
    try {
        const {
            id
        } = req.params;
        await deleteEvent(parseInt(id));
        res.sendStatus(204);
    } catch (error) {
        res.status(500).send(error);
    }
});

// Ruta para crear un nuevo usuario
app.post('/users/:id', async (req: Request, res: Response) => {
    try {
        const {
            id
        } = req.params;
        const result = await createUser(parseInt(id));
        res.json({
            id
        });
    } catch (error) {
        res.status(500).send(error);
    }
});

// Ruta para eliminar usuario
app.delete('/users/:id', async (req: Request, res: Response) => {
    try {
        const {
            id
        } = req.params;
        const result = await deleteUser(parseInt(id));
        res.json({
            id
        });
    } catch (error) {
        res.status(500).send(error);
    }
});

// Iniciar el servidor en el puerto 3000
app.listen(3000, () => {
    console.log("Servidor web iniciado en el puerto 3000");
});

// Función para obtener todos los eventos
function getAllEvents() {
    return new Promise((resolve, reject) => {
        const sql = 'SELECT * FROM events';
        db.query(sql, (err, result) => {
            if (err) {
                console.error('Error al obtener eventos: ', err);
                return reject(err);
            }
            return resolve(result);
        });
    });
}

// Función para obtener todos los eventos de un usuario
function getAllEventsByUserId(userId: number) {
    return new Promise((resolve, reject) => {
        const sql = `SELECT * FROM events JOIN (SELECT eventId FROM participants WHERE userId = ${userId}) AS usr ON id = usr.eventId`;
        db.query(sql, (err, result) => {
            if (err) {
                console.error('Error al obtener eventos: ', err);
                return reject(err);
            }
            return resolve(result);
        });
    });
}

// Función para crear un nuevo evento
function createEvent(title: string, description: string, start: string, end: string, allDay: boolean, location: string, userId: number) {
    return new Promise((resolve, reject) => {
        const sql = `INSERT INTO events (title, description, start, end, allDay, location) VALUES ('${title}', '${description}', '${start}', '${end}', ${allDay}, '${location}')`;
        db.query(sql, (err, result) => {
            if (err) {
                console.error('Error al crear evento: ', err);
                return reject(err);
            }
            console.log('Evento creado con éxito');
            const json: any = result;
            const evenId: number = json.insertId;
            const result_2 = createParticipant(evenId, userId);
            return resolve(result);
        });
    });
}

// Función para crear un nuevo usuario
function createUser(id: number) {
    return new Promise((resolve, reject) => {
        const sql = `INSERT INTO users (id) VALUES ('${id}')`;
        db.query(sql, (err, result) => {
            if (err) {
                console.error('Error al crear usuario: ', err);
                return reject(err);
            }
            console.log('Usuario creado con éxito');
            const json: any = result;
            return resolve(result);
        });
    });
}

// Función para eliminar un usuario existente
function deleteUser(id: number) {
    return new Promise((resolve, reject) => {
        const sql = `DELETE FROM users WHERE id=${id}`;
        db.query(sql, (err, result) => {
            if (err) {
                console.error('Error al eliminar udustio: ', err);
                return reject(err);
            }
            console.log('Usuario eliminado con éxito');
            return resolve(result);
        });
    });
}

//FUNCION PARA AGREGAR PARTICIPANTES
function createParticipant(eventId: number, userId: number) {
    return new Promise((resolve, reject) => {
        const sql = `INSERT INTO participants (eventId,userId) VALUES ( '${eventId}','${userId}')`;
        db.query(sql, (err, result) => {
            if (err) {
                console.error('Error al agregar participante: ', err);
                return reject(err);
            }
            console.log('Participante creado con éxito');
            return resolve(result);
        });
    });
}

// Función para actualizar un evento existente
function updateEvent(id: number, title: string, description: string, start: string, end: string, allDay: boolean, location: string) {
    return new Promise((resolve, reject) => {
        const sql = `UPDATE events SET title='${title}', description='${description}', start='${start}', end='${end}', allDay=${allDay}, location='${location}' WHERE id=${id}`;
        db.query(sql, (err, result) => {
            if (err) {
                console.error('Error al actualizar evento: ', err);
                return reject(err);
            }

            console.log('Evento actualizado con éxito');
            return resolve(result);
        });
    });
}

// Función para eliminar un evento existente
function deleteEvent(id: number) {
    return new Promise((resolve, reject) => {
        const result_2 = deleteParticipant(id)
        const sql = `DELETE FROM events WHERE id=${id}`;
        db.query(sql, (err, result) => {
            if (err) {
                console.error('Error al eliminar evento: ', err);
                return reject(err);
            }
            console.log('Evento eliminado con éxito');
            return resolve(result);
        });
    });
}

function deleteParticipant(id: number) {
    return new Promise((resolve, reject) => {
        const sql = `DELETE FROM participants WHERE eventId=${id}`;
        db.query(sql, (err, result) => {
            if (err) {
                console.error('Error al eliminar participantes asociados: ', err);
                return reject(err);
            }
            console.log('participantes asociados eliminados con éxito');
            return resolve(result);
        });
    });
}


//mqGetEventsClient()
//mqCreateEventClient()