import React from 'react'
import {View, Button} from 'react-native'

export default class LoginScreen extends React.Component {
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
      fetch('http://192.168.0.3:5000/')
        .then(res=>console.log(res))
        .catch(err=>console.log(err))
    };
  }
  

  