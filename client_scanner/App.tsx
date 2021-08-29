/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * Generated with the TypeScript template
 * https://github.com/react-native-community/react-native-template-typescript
 *
 * @format
 */

 'use strict';

import React from 'react';
import {
  SafeAreaView,
  ScrollView,
  StatusBar,
  Text,
  useColorScheme,
  View,
} from 'react-native';
import { BarCodeReadEvent } from 'react-native-camera';
import QRCodeScanner from 'react-native-qrcode-scanner';

import {
  Colors,
  Header
} from 'react-native/Libraries/NewAppScreen';

const App = () => {
  const isDarkMode = useColorScheme() === 'dark';

  const backgroundStyle = {
    backgroundColor: isDarkMode ? Colors.darker : Colors.lighter,
  };

  let scanner: QRCodeScanner;

  const onCodeScanned = (e: BarCodeReadEvent) => {
    scanner.reactivate()
    console.log(`Found barcode: ${e.data}`)
  }

  return (
    <SafeAreaView style={backgroundStyle}>
      <StatusBar barStyle={isDarkMode ? 'light-content' : 'dark-content'} />
      <ScrollView
        contentInsetAdjustmentBehavior="automatic"
        style={backgroundStyle}>
        <Header />
        <View
          style={{
            backgroundColor: isDarkMode ? Colors.black : Colors.white,
          }}>
          <QRCodeScanner
            ref={(node) => { scanner = node; }}
            onRead={onCodeScanned}
            topContent={
              <Text>
                Scan some code
              </Text>
            }
            bottomContent={
              <Text>hello</Text>
            }
        />
        </View>
      </ScrollView>
    </SafeAreaView>
  );
};

export default App;
