import React from 'react'

export class Home extends React.Component{
    render(){
        const { name } = this.props.user
        return (
            <div>welcome home, {name}</div>
        )
    }
}