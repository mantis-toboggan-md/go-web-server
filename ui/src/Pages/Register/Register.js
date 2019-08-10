import React from 'react'

export class Register extends React.Component{

    render(){
        return(
            <div>
                <h3>register</h3>
                name:
                <input type = 'text' id='name'/> <br/>
                password:
                <input type = 'text' id='password'/> <br/>
                re-enter password:
                <input type = 'text' id='passwordConfirm'/> <br/>
             </div>   
        )
    }
}