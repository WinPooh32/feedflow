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
    passwordValid: boolean
 }

export class Login extends React.Component<PropsState, LoginState> {
    private api = new PagesApi()

    constructor(props: PropsState, state: LoginState){
        super(props, state)

        this.state = {
            form: {username: '', password: ''},
            usernameValid: true,
            passwordValid: true
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
        const value = target.type === 'checkbox' ? target.checked : target.value;
        const name = target.name;

        this.setState((state): LoginState => {
            state[name] = value
            return state
        });

        if (name === "username") {
            this.validateName(value as string)
        } else if (name == "password"){
            this.validatePassword(value as string)
        }
    }
    
    handleSubmit(event: React.FormEvent) {
        event.preventDefault();

        let state = this.state
        if (state.usernameValid && state.passwordValid){
            this.api.login(this.state.form).then(resp => {
                let status = resp.status
                if(status !== 200){return}
    
                console.log(status, "Logged in!!!")
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
                            placeholder="Password" 
                            autoCapitalize="none" 
                            autoCorrect="off" 
                            autoComplete="password" 
                            minLength={minPwdLength} maxLength={maxPwdLength}
                            required={false} 
                        />
                        <label htmlFor="inputPassword">Пароль</label>
                    </div>

                    <div className="form-label-group">
                        <button className="btn btn-lg btn-success btn-block" type="submit">Войти</button>                                      
                    </div>

                    <div className="m-auto text-center">
                        <a href="/signin">РЕГИСТРАЦИЯ</a>
                    </div>
                </form>
            </div>
        );
    }
}
