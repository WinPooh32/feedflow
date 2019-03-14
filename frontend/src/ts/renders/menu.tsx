import { Menu } from "../components/menu/Menu";
import React = require("react");
import ReactDOM = require("react-dom");

export function Render(){
    let node = document.getElementById("menu")

    //Read attributes
    let loggedId = node.getAttribute("loggedIn") === "true"
    let admin = node.getAttribute("admin") === "true"

    ReactDOM.render(
        <Menu loggedIn={loggedId} admin={admin} />,
        node
    );
}