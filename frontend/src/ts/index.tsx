export * from "./api";
export * from "./configuration";

import * as React from "react"
import * as ReactDOM from "react-dom"
import { FeedApi, PagesApi, NewPageContent } from "./api";

const feedApi = new FeedApi()
const pagesApi = new PagesApi()

import { Hello } from "./components/Hello";

ReactDOM.render(
    <Hello compiler="TypeScript" framework="React" />,
    document.getElementById("example")
);


