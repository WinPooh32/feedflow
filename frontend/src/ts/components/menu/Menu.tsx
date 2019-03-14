import * as React from "react";
import {Login} from "./Login";

export interface MenuProps { loggedIn: boolean; admin: boolean; }
export interface MenuState { menu: Array<React.ReactElement> }

export class Menu extends React.Component<MenuProps, MenuState> {
    
    constructor(props: MenuProps){
        super(props)
        this.state = {menu: new Array<React.ReactElement>()} as MenuState
    }

    render() {
        let menu = this.state.menu

        if(this.props.loggedIn){
            if(this.props.admin){

            }
        }else{
            menu.push(<Login key="login"/>)
        }

        return menu
    }
}


