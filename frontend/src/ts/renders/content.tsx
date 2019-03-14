import { Hello } from "../components/Hello";
import ReactDOM = require("react-dom");
import React = require("react");

export function Render(){
    ReactDOM.render(
        <Hello compiler="TypeScript" framework="React" />,
        document.getElementById("example")
    );
}