// let conn;
// let selectedChat = "general";

// const changeChatRoom = () => {
//     const newchat = document.getElementById('chatroom');

//     if(newchat != null && newchat.value != selectedChat) {
//         console.log(newchat)
//     }
//     return false; 
// }

// const sendMessage = () => {
//     const newmessage = document.getElementById('message');

//     if( newmessage != null ) {
//         conn.send(newmessage.value)
//     }
//     return false;
// }


// window.onload = () => {
//     document.getElementById('chatrom-selection').onsubmit = changeChatRoom();
//     document.getElementById('chatroom-message').onsubmit = sendMessage();

//     if(window["WebSocket"]) {
//         console.log("supports websockets");
//         // this make a pettition to the endpoint from our api
// /* 
//         TODO: para el futuro desarrollo de esta onda, existen dos protocolos para conectarse,
//         el ws:// y el wss://, ahorita estamos trabajando con ws:// y este protocolo es de websockets no seguros, para 
//         desarrollar ahorita y sacar la demo, esta bien, pero para desarrollo tenemos que ver como manejar el wss
// */
//         conn = new WebSocket("ws://" + document.location.host + "/ws");
//     } else {
//         alert('Browser does not support websockets');
//     }
// }


if(window["WebSocket"]) {
    console.log("supports websockets");
    var conn =  new WebSocket("ws://" + document.location.host + "/ws");
    conn.send("hola")
} else {
    alert('Browser does not support websockets');
}