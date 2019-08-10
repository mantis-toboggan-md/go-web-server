import React from 'react'

export default class RegisterScreen extends React.Component {
    static navigationOptions = {
      title: 'Please sign in',
    };
  
    render() {
      return (
        <View>
          <Button title="Sign in!" onPress={this._signInAsync} />
        </View>
      );
    }
  
    _signInAsync = async () => {
      fetch('localhost:5000/')
        .then(res=>console.log(res))
    };
  }
  