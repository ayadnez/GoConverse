// app/index.js 

var socket = new WebSocket("ws://localhost:8081/ws");

let connect  = () => {
    console.log("Attempting connection...");

    socket.onopen = () => {
        console.log("successfully connected");
    };

    socket.onmessage = msg => {
        console.log(msg);
    };

    socket.onclose = event => {
        console.log("socket closed connection : ", event);
    };

    socket.onerror = error => {
        console.log("socket Error: ", error);
    };
};

let sendMsg = msg => {
    console.log("socket Error: ", msg);
    socket.send(msg);
};

export {connect,sendMsg};