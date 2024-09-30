import React, { Component } from 'react';
import {AppRegistry, Window, App} from 'proton-native';



class win extends Component {
    render() {
      return (
        <App>
          <Window style={{ height: '50%', width: '20%' }} />
        </App>
      );
    }
}

AppRegistry.registerComponent('Test', <Example />);


const main = () => {
    const win = new win;
    win.render();
}