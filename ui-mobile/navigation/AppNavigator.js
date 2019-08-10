import React from 'react';
import { createAppContainer, createSwitchNavigator, createStackNavigator } from 'react-navigation';
import { Platform } from 'react-native';
import MainTabNavigator from './MainTabNavigator';
import AuthLoadingScreen from '../screens/AuthLoadingScreen'
import LoginScreen from '../screens/LoginScreen'
import RegisterScreen from '../screens/RegisterScreen'




const config = Platform.select({
  web: { headerMode: 'screen' },
  default: {},
});

const AuthStack = createStackNavigator(
  {
    Login: LoginScreen,
    Register: RegisterScreen
  },
  config
);

AuthStack.navigationOptions = {
  tabBarLabel: 'Log in',
  tabBarIcon: ({ focused }) => (
    <TabBarIcon
      focused={focused}
      name={
        Platform.OS === 'ios'
          ? `ios-information-circle${focused ? '' : '-outline'}`
          : 'md-information-circle'
      }
    />
  ),
};

AuthStack.path = '';

export default createAppContainer(
  createSwitchNavigator({
    // You could add another route here for authentication.
    // Read more at https://reactnavigation.org/docs/en/auth-flow.html
    Main: MainTabNavigator,
    AuthLoading: AuthLoadingScreen,
    Auth: AuthStack,   
  },
  {
    initialRouteName: 'AuthLoading',
  })
);