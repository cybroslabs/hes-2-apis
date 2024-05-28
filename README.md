# hes-2-apis

- Modules
  - openapi
  - protobuf

---
Q/As:
- Jak mají vypadat EPs? Příklad: Uvažujeme, že máme job-bulks. Mají existovat v API EPs pro DataProxy samostatně?
  - Ne.
  * StartSingleJob
    - nemá tohle být buď list-of-jobs a nebo bulk o jednom záznamu?
  * StartListOfJobs <- does this make sense? Ano; mixed-device groups. Ale neni to jeden bulk.
    - nemá tohle být list-of-bulks?
  * StartBulkJob <- many-devices, shared settings & actions
    - nemá tohle být list-of-bulks?

    ^^^- všechno výše by mělo umět device_id bez dev params, který by si to mělo umět dotahat z registered devices.

- Co potřebujeme o jobech?
  - GetJobs <- job listing, co to má vracet?
      ... má to vracet virtuální bulk-id pro single-joby? viz výše
  - GetJobInfo <- co to má vracet?
  - GetData ... uvažujeme-li raw tak by tohle mělo být single-job (filter == bulkid & deviceid)

- Co potřebujem na device-list?
  - RegisterDevice <- new device (device_id, connection_info)
    ... co když se zavolá job a device není registered? má se auto-registrovat z info ve start-bulk?
  - UnregisterDevice <- remove device    
  