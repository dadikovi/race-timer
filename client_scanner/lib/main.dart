import 'package:flutter/material.dart';
import './qr_scanner_state.dart';

void main() {
  runApp(ScannerApp());
}

class ScannerApp extends StatelessWidget {

  static const APP_NAME = 'Race-timer Scanner';

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: APP_NAME,
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: ScannerHome(title: APP_NAME),
    );
  }
}

class ScannerHome extends StatefulWidget {
  ScannerHome({Key key, this.title}) : super(key: key);

  final String title;

  @override
  State<StatefulWidget> createState() => QRScannerState();
}