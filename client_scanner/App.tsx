'use strict';

import React from 'react';
import { NativeSyntheticEvent, ToastAndroid } from 'react-native';
import { NativeTouchEvent } from 'react-native';
import { Button } from 'react-native';
import {
  SafeAreaView,
  ScrollView,
  StatusBar,
  Text,
  TextInput,
  View,
  Switch
} from 'react-native';
const axios = require('axios').default;
import { Card } from 'react-native-elements'
import { BarCodeReadEvent } from 'react-native-camera';
import QRCodeScanner from 'react-native-qrcode-scanner';
import { useEffect } from 'react';
//import { Switch } from 'react-native-elements/dist/switch/switch';

const App = () => { 

  let scanner: QRCodeScanner;

  const [finishMode, onFinishModeChange] = React.useState(true);
  const [serverCode, onServerCodeChange] = React.useState('');
  const [serverConnectionOk, setServerConnectionOk] = React.useState(false);

  useEffect(() => {
    setServerConnectionOk(false)
  }, [serverCode])

  const onCodeScanned = (e: BarCodeReadEvent) => {
    if (finishMode) {
      axios.post(serverUrl(`participants/${e.data}`))
      .then((r:any) => {
        scanner.reactivate()
        console.log(`Finished participant: ${e.data}`)
      })
      .catch((r:any) => {
        scanner.reactivate()
        ToastAndroid.showWithGravity(
          `Error, captured data: ${e.data}`,
          ToastAndroid.LONG,
          ToastAndroid.BOTTOM
        );
      })
    } else {
      axios.post(serverUrl(`participants`), {startNumber: new Number(e.data)})
        .then((r:any) => {
          scanner.reactivate()
          console.log(`Finished participant: ${e.data}`)
        })
        .catch((r:any) => {
          scanner.reactivate()
          ToastAndroid.showWithGravity(
            `Error, captured data: ${e.data}`,
            ToastAndroid.LONG,
            ToastAndroid.BOTTOM
          );
        })
    }
  }

  const syncWithServer = (e: NativeSyntheticEvent<NativeTouchEvent>) => {
    axios.get(serverUrl('segments'))
      .then((r:any) => setServerConnectionOk(true))
      .catch((r:any) => {
        setServerConnectionOk(false)
        ToastAndroid.showWithGravity(
          `Error: ${r}`,
          ToastAndroid.LONG,
          ToastAndroid.BOTTOM
        );
      })
  }

  const serverUrl = (path: string) => {
    return `http://${serverCode}.ngrok.io/${path}`
  }

  return (
    <SafeAreaView>
      <StatusBar />
      <ScrollView contentInsetAdjustmentBehavior="automatic">
        <Card containerStyle = {{borderColor: serverConnectionOk ? 'green' : 'red'}}>
          <View style = {{flexDirection: "row"}}>    
            <View style={{ flex: 4}}>
              <TextInput 
                placeholder={'Server code'}
                onChangeText={onServerCodeChange}
                value={serverCode}
              />
            </View>
            <View style={{ flex: 1}}>
              <Button 
                title={'OK'}
                onPress={syncWithServer}
              />
            </View>
          </View>
        </Card>
        <View>
          <QRCodeScanner
              ref={(node) => { scanner = node; }}
              onRead={onCodeScanned}
            />
        </View>
        <Card>
        <View style = {{flexDirection: "row"}}>    
            <View style={{ flex: 4}}>
              <Text>Finish mode?</Text>
            </View>
            <View style={{ flex: 1}}>
              <Switch 
                value={finishMode} 
                thumbColor={finishMode ? "blue" : "blue"}
                onValueChange={(switchValue) => onFinishModeChange(switchValue)}
              />
            </View>
          </View>
        </Card>
      </ScrollView>
    </SafeAreaView>
  );
};

export default App;
