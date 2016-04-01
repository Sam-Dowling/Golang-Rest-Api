
from flask import Flask, Response, request
from urllib import unquote
import json
from FleetAPI import FleetAPI, prettify


app = Flask(__name__)


def create_dict(data):
    request_dict = {}
    item_list = data.split("&")
    for item in item_list:
        split_item = item.split("=")
        request_dict[split_item[0]] = unquote(split_item[1]).replace("+", " ")
    return request_dict


@app.route("/", methods=['POST', 'GET'])
def main():
    api = FleetAPI("sam@email.com", "password", '127.0.0.1')
    data = request.get_data()
    request_dict = create_dict(data)
    carriercode = request_dict['text'][6:]
    payload = prettify(api.get_fleet_data(carriercode)[-1])
    json_response = json.dumps({"text": "Hi %s\nThe most recent data I have for %s is:\n%s" % (
        request_dict['user_name'], carriercode, payload)})
    response = Response(json_response, status=200, mimetype='application/json')
    return response


if __name__ == "__main__":
    app.run('0.0.0.0', 8000, debug=True)
