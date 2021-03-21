import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class CreateSegment extends StatefulWidget {
  final ValueChanged<String> onCreated;

  const CreateSegment({Key key, this.onCreated}) : super(key: key);

  @override
  State<StatefulWidget> createState() {
    return new CreateSegmentState(onCreated);
  }
}

class CreateSegmentState extends State<CreateSegment> {
  final TextEditingController _controller = TextEditingController();
  final ValueChanged<String> onCreated;

  CreateSegmentState(this.onCreated);

  @override
  Widget build(BuildContext context) {
    return Row(children: [
      TextField(
        controller: _controller,
        onSubmitted: (String value) async {
          var response = await http.post(
              Uri.parse('http://localhost:8010/segments'),
              body: {'name': value});

          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(content: Text('Response body: ${response.body}')),
          );
          onCreated.call(value);
        },
        decoration: InputDecoration(
          border: OutlineInputBorder(),
          labelText: 'New segment name',
        ),
      )
    ]);
  }
}
