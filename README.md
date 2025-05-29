# expfake

## example usage

```bash
 $ go run expfake.go --filesizes filesizes.yaml --hostmap hostmap.yaml --hostname lsstcam-dc03 --dir /tmp/test1
2025/05/29 10:28:48 Generated file /tmp/test1/R12_S00.fits with size 27342720 bytes
2025/05/29 10:28:48 Generated file /tmp/test1/R12_S00.json with size 3054 bytes
2025/05/29 10:28:48 Generated file /tmp/test1/R12_S01.fits with size 26504640 bytes
2025/05/29 10:28:48 Generated file /tmp/test1/R12_S01.json with size 3054 bytes
2025/05/29 10:28:48 Generated file /tmp/test1/R12_S02.fits with size 27394560 bytes
2025/05/29 10:28:48 Generated file /tmp/test1/R12_S02.json with size 3054 bytes
2025/05/29 10:28:48 Generated file /tmp/test1/R12_S10.fits with size 27120960 bytes
2025/05/29 10:28:48 Generated file /tmp/test1/R12_S10.json with size 3053 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R12_S11.fits with size 26890560 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R12_S11.json with size 3053 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R12_S12.fits with size 27060480 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R12_S12.json with size 3053 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R12_S20.fits with size 27383040 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R12_S20.json with size 3053 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R12_S21.fits with size 27227520 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R12_S21.json with size 3053 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R12_S22.fits with size 27011520 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R12_S22.json with size 3053 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S00.fits with size 26752320 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S00.json with size 3053 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S01.fits with size 26660160 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S01.json with size 3053 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S02.fits with size 26982720 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S02.json with size 3053 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S10.fits with size 41457600 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S10.json with size 3053 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S11.fits with size 41679360 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S11.json with size 3053 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S12.fits with size 40968000 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S12.json with size 3053 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S20.fits with size 26781120 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S20.json with size 3053 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S21.fits with size 26951040 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S21.json with size 3053 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S22.fits with size 26809920 bytes
2025/05/29 10:28:49 Generated file /tmp/test1/R22_S22.json with size 3053 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S00.fits with size 28036800 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S00.json with size 3054 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S01.fits with size 27100800 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S01.json with size 3054 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S02.fits with size 26533440 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S02.json with size 3054 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S10.fits with size 26640000 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S10.json with size 3053 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S11.fits with size 27077760 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S11.json with size 3053 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S12.fits with size 26691840 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S12.json with size 3053 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S20.fits with size 27273600 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S20.json with size 3054 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S21.fits with size 27411840 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S21.json with size 3054 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S22.fits with size 27144000 bytes
2025/05/29 10:28:50 Generated file /tmp/test1/R34_S22.json with size 3054 bytes
2025/05/29 10:28:50 Total size of files generated for host lsstcam-dc03: 772970760 bytes
```
