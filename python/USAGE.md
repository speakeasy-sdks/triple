<!-- Start SDK Example Usage -->


```python
import triple


s = triple.Triple()


res = s.pets.create_pets()

if res.status_code == 200:
    # handle response
```
<!-- End SDK Example Usage -->