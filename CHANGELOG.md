# Changelog

## 0.1.0-alpha.3 (2025-08-22)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/Metronome-Industries/terraform-provider-metronome/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** api update ([3abaec4](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/3abaec457971f14bd6065ff7c4bca1a6284cefe7))
* **api:** api update ([471c750](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/471c750e49acbe5c916a547669a449da4f30104a))
* ensure `internal/apiform` encoder can handle "force_encode" serialization tag ([bb6dfdd](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/bb6dfdd573b1b79d81ce127411e0634b2deb3001))
* new option to send computed values back to server ([d340374](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/d3403742a77efbfe98cdae3f5f9259d85a8103b6))


### Bug Fixes

* **api:** handle mismatched dynamic array types in state and plan during serialization ([3df79a0](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/3df79a04a9ac98327ae1a8a93d0a019f44c4ffd8))
* **ci:** release-doctor — report correct token name ([331f6dc](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/331f6dc23fad047aca9f789e75f6cff819efd2a2))
* dynamic type validators should handle int and floats correctly ([f459e17](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/f459e17f35c30248fb4912f83ec1b876ac8a6d5f))
* encoder crash for nested nils in dynamic types ([7603da4](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/7603da41fa42e5b4387afdc9828d89c9ade751bf))
* null nested attribute decoding ([c7b475b](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/c7b475b7e9037eedd1c742723c69865b1ded8015))
* populate computed_optional collections from API responses ([f3f117d](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/f3f117d27aa2062ec8324e53b8c63a8a3ca729ba))


### Chores

* **ci:** enable for pull requests ([ac043bb](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/ac043bb0107a5958ba520386fc5e58039b7a5dfb))
* **ci:** only run for pushes and fork pull requests ([8a6f96f](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/8a6f96f573b05583ca14f5702d0c4a3663c6cf9b))
* **internal:** codegen related update ([af1938f](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/af1938f1d6322ba8e4491d39a5a8d91925c9a0f5))
* **internal:** codegen related update ([04f57a5](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/04f57a5846de37919f64bce684c64736795ccd1f))
* **internal:** upgrade cloudflare/circl ([5d33889](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/5d33889d1fbee8d0b10e215edc0ce300eeb298db))
* update @stainless-api/prism-cli to v5.15.0 ([2486297](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/2486297fad4f421ba83ef98745cc8967e26e8f01))
* update SDK settings ([c31dcbe](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/c31dcbec9e4c0e0896edf15fb24dace7c684b82a))

## 0.1.0-alpha.2 (2025-06-09)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/Metronome-Industries/terraform-provider-metronome/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** infer all services ([1445ff5](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/1445ff50e08468b0140a3f9fb7ae382499fa0cec))

## 0.1.0-alpha.1 (2025-06-02)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/Metronome-Industries/terraform-provider-metronome/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** api update ([af15a2e](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/af15a2e4ad62763dd077e7c77d55d4f4b80cf775))
* **api:** update via SDK Studio ([bc7a0d2](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/bc7a0d2261a2737ecacedc5377e1476584446e17))
* **api:** update via SDK Studio ([24b6cd2](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/24b6cd234632d6f00b76e42dec7166809e4eccfb))
* **client:** support environments property from Stainless config ([61045df](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/61045dfbab6a9d090312abf535a6c2f09e27b200))


### Bug Fixes

* **build:** enable building against private Go production repos ([8d91c67](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/8d91c671fadc5f13e7894d4b2cbb31487fc44ecd))


### Chores

* bump deps to avoid GetResourceIdentitySchemas errors for Terraform CLI v1.12+ ([50854ca](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/50854cabe850cd51ff33407087951263e1c3dd3c))
* configure new SDK language ([78208ca](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/78208cac02cef64f3fe66f754bf526278a17e81b))
* configure new SDK language ([274e522](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/274e5224fe74887e75e153e366d9defa0df37502))
* **docs:** grammar improvements ([4b63400](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/4b634003635bcee03e530e514972e301895d2064))
* **internal:** codegen related update ([72d52ec](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/72d52eca0efe9d06fafb171c6384c946c9742d01))
* **internal:** codegen related update ([b01ab0a](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/b01ab0a2304168b145ea4b27ca550831f90977a5))
* update SDK settings ([516907b](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/516907b78611bbd6234b884eab60b2a092d117d6))
* update SDK settings ([27f1bde](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/27f1bdee1111fc1565cffa524d1c7185d073ac38))
