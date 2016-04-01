import json
import urllib3
urllib3.disable_warnings()

def prettify(message):
    return "\nCarrier: %s\nCount: %s\n\n" % (message['carriercode'], message['aircraft_count'])

class FleetAPI:
    def __init__(self, email, password, ip):
        self.ip = ip
        self.pool = self.create_pool('FleetRest/key/server-local.cert')
        self.sessionToken = self.login(email, password)


    def create_pool(self, cert_location):
        return urllib3.PoolManager(
            cert_reqs='CERT_REQUIRED',
            ca_certs=cert_location
        )


    def login(self, email, password):
        url = 'https://%s:8080/login' % self.ip
        headers = {'content-type': 'application/json'}
        data = '{"email": "%s", "password": "%s"}' % (email, password)
        try:
            res = self.pool.urlopen('PUT', url, headers=headers, body=data)
            if res.status == 200:
                return json.loads(res.data.decode(encoding='UTF-8'))['sessionToken']
            else:
                return None
        except urllib3.exceptions.SSLError:
            return None


    def get_fleet_data(self, carrier_code):
        url = 'https://%s:8080/%s' % (self.ip, carrier_code)
        headers = {'Authorization' : 'Bearer %s' % self.sessionToken}
        r = self.pool.request('GET', url, headers=headers)
        return json.loads(r.data.decode(encoding='UTF-8'))
