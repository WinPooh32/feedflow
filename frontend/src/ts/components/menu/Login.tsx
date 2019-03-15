import * as React from "react";
import { PagesApi, LoginRequest } from "../../api";

const minNameLength = 2
const maxNameLength = 32

const minPwdLength = 10
const maxPwdLength = 32

export interface PropsState {}
export interface LoginState { 
    form: {username: string, password: string},
    usernameValid: boolean,
    passwordValid: boolean,
    loading: boolean
 }

export class Login extends React.Component<PropsState, LoginState> {
    private api = new PagesApi()

    constructor(props: PropsState, state: LoginState){
        super(props, state)

        this.state = {
            form: {username: '', password: ''},
            usernameValid: true,
            passwordValid: true,
            loading: false
        } as LoginState
        
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    private set usernameValid(v: boolean){
        this.setState(() => {
            return {usernameValid: v}
        });
    }

    private set passwordValid(v: boolean){
        this.setState(() => {
            return {passwordValid: v}
        });
    }

    private validateName(name: string){
        if(name.length < minNameLength || name.length > maxNameLength) { 
            this.usernameValid = false
        }else{
            this.usernameValid = true
        }
    }

    private validatePassword(pwd: string){
        if(pwd.length < minPwdLength || pwd.length > maxPwdLength) { 
            this.passwordValid = false
        }else{
            this.passwordValid = true
        }
    }

    private handleChange(event: React.ChangeEvent<HTMLInputElement>) {
        const target = event.target;
        const value = target.value;
        const trigger = target.type === 'checkbox' ? target.checked: false
        const name = target.name;

        this.setState((state): LoginState => {
            let mystate;
            switch(name){
                case "username": 
                case "email":
                case "password":{
                    mystate = state.form
                    break
                }

                default:{
                    mystate = state
                }
            }

            mystate[name] = value

            return state
        });

        if (name === "username") {
            
        } else if (name == "password"){
            
        }
    }

    disableLoading(){
        setTimeout(()=>{
            this.setState({loading: false})
        }, 2 * 1000)
    }
    
    handleSubmit(event: React.FormEvent) {
        event.preventDefault();

        let state = this.state
        if (state.usernameValid && state.passwordValid && !state.loading ){

            this.setState({loading: true})

            this.api.login(this.state.form).then(resp => {
                let status = resp.status
                if(status !== 200){return}
    
                console.log(status, "Logged in!!!")
                this.disableLoading()
            }).catch(() => {
                this.disableLoading()
            })
        }
    }

    render() {
        let validNameClass = this.state.usernameValid ? "is-valid" : "is-invalid"
        let validPwdClass = this.state.passwordValid ? "is-valid" : "is-invalid"

        return (
            <div className="menu-block-wrap">
                <form onSubmit={this.handleSubmit} className="menu-block form-signin m-auto needs-validation">
                    <input type="hidden" autoFocus/>
                    <div className="form-label-group has-error has-feedback form-group">
                        <input 
                            defaultValue={this.state.form.username} 
                            onChange={this.handleChange} 
                            type="username" id="username" name="username" 
                            autoCapitalize="none" 
                            autoCorrect="off" 
                            autoComplete="username" 
                            minLength={minNameLength} maxLength={maxNameLength}
                            className={"form-control " + validNameClass } 
                            placeholder="Имя пользователя" 
                            required={false}
                        />
                        <label htmlFor="username">Имя пользователя</label>
                    </div>
                    <div className="form-label-group has-error">
                        <input 
                            defaultValue={this.state.form.password} 
                            onChange={this.handleChange} 
                            type="password" 
                            id="inputPassword" 
                            name="password" 
                            className={"form-control " + validPwdClass}
                            placeholder="Пароль" 
                            autoCapitalize="none" 
                            autoCorrect="off" 
                            autoComplete="password" 
                            minLength={minPwdLength} maxLength={maxPwdLength}
                            required={false} 
                        />
                        <label htmlFor="inputPassword">Пароль</label>
                    </div>

                    <div className="form-label-group">
                        <button className="btn btn-lg btn-success btn-block" type="submit"><CompButtonContent showLoader={this.state.loading}/></button>                                      
                    </div>

                    <div className="m-auto text-center">
                        <a href="/signin">РЕГИСТРАЦИЯ</a>
                    </div>
                </form>
            </div>
        );
    }
}

export interface ButtonContentProps {
    showLoader: boolean
}

function CompButtonContent(props: ButtonContentProps): React.ReactElement<any>{
    if(props.showLoader){
        return <div className="loader m-auto"></div>
    }else{
        return <span>Войти</span>
    }
}