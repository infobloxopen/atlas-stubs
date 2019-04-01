# authn-stub

The authn-stub represents the stub for authentication service and serves
solely for multi-accounting demo purposes. This service handles incoming GET
requests on /auth URL if specified header `Throw-Token` returns `Grpc-Metadata-Authorization` header with value from `Throw-Token` if not depending on `User-And-Pass`  header contents returns
either 401 (UNAUTHORIZED) or 200 (OK) with `Grpc-Metadata-Authorization` header
set.

Following User-And-Pass values are accepted:

admin1:admin -- for AccountID=1
admin2:admin -- for AccountID=2



## NGINX deployment notes

To set-up AuthN chain in a k8s cluster following steps should be completed:

1. Deploy AuthN (or AuthN-Stub) using manifest in deploy/ folder
2. Add annotations to a service that is a successor of an authentication process:

```
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  namespace: XXX
  name: YYY
  annotations:
    ingress.kubernetes.io/auth-url: http://authn-stub.atlas-stubs.svc.cluster.local/auth
    ingress.kubernetes.io/auth-response-headers: Grpc-Metadata-Authorization
```
