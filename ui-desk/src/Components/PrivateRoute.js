import React from 'react'
import { Redirect } from 'react-router-dom';
/**
 * privteRoute wraps react-router routes to check /auth and redirect to login
 * renders 'authenticating...' while checking auth
 * passes route to login to redirect back to desired page
 * @param {Component} Protected - component to hide behind auth
*/


export function privateRoute(Protected){
    return class extends React.Component{
        state = {
            loading: true,
            authenticated: false
        }

        // check token validity
        componentDidMount(){
            fetch("/auth", {
                method: 'GET',
                withCredentials: true,
                credentials: 'include',
                headers: {
                    authorization: localStorage.getItem("token")
                }
            })
            .then(res=>{
                if(res.ok){
                    this.setState({loading: false, authenticated: true})
                } else {
                    const error = new Error(res.error);
                    throw error;  
                }         
            })
            .catch(()=>this.setState({loading:false, authenticated: false}))
        }

        render(){
            const { loading, authenticated } = this.state
            // if loading completed and not authenticated, token was bad
            const redirect = !loading && !authenticated
            return redirect
                // pass along location to go back to the protected page after login
                ? <Redirect to={{
                    pathname: '/login',
                    state: { from: this.props.location }
                  }} /> 
                : loading ?
                    <div>authenticating...</div>
                    // final condition of ternary: not loading and authenticated
                    : ( <React.Fragment>
                            <Protected {...this.props} />
                        </React.Fragment> )
        }
    }
}

