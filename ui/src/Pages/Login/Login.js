import React from 'react'

export class Login extends React.Component{

    render(){
        return(
            <div>
                <h3>login</h3>
                name:
                <input type = 'text' id='name'/> <br/>
                password:
                <input type = 'text' id='password'/> <br/>
             </div>   
        )
    }
}