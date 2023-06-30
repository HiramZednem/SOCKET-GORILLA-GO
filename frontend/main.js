let selectedChat = "general";

const changeChatRoom = () => {
    const newchat = document.getElementById('chatroom');

    if(newchat != null && newchat.value != selectedChat) {
        console.log(newchat)
    }
    return false;
}

const sendMessage = () => {
    const newmessage = document.getElementById('message');

    if( newmessage != null ) {
        console.log(newmessage)
    }
    return false;
}


window.onload = () => {
    document.getElementById('chatrom-selection').onsubmit = changeChatRoom();
    document.getElementById('chatroom-message').onsubmit = sendMessage();

    if(window["WebSocket"]) {
        console.log("supports websockets");
        // this make a pettition to the endpoint from our api
        let conn = new WebSocket("ws://" + document.location.host + "/ws");
    } else {
        alert('Browser does not support websockets');
    }
}