import { Hello } from "../components/Hello";
import ReactDOM = require("react-dom");
import React = require("react");
import { Signon } from "../components/menu/Signon";

export function Render(){
    ReactDOM.render(
        <Signon/>,
        document.getElementById("signonform")
    );
}