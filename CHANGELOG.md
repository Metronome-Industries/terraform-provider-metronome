# Changelog

## 0.1.0-alpha.4 (2026-07-21)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/Metronome-Industries/terraform-provider-metronome/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* [ORCH-1875] add billing config schedule to api ([fa955df](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/fa955df27c1cc71287989ba6285c1f868620fbbb))
* add stlc SDK generation workflow to api repo ([6b9c983](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/6b9c98306f7b61d697f189d6059368be0b18d1fb))
* Show Tax/PDF Fields on Metronome Invoice ([9b2cd14](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/9b2cd14bbc37944b49ee69e598c4354e80dfecfb))
* **stlc:** configurable CI runner and private-production-repo support in workflow templates ([3c50ee3](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/3c50ee302a94bade3baca8822fc1069641906e52))


### Bug Fixes

* **terraform:** fix nil pointer dereference when comparing null or unknown numbers ([486d43f](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/486d43f34bebe8c8f7abfd3f61b3f2487ac03fb7))

## 0.1.0-alpha.3 (2026-06-01)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/Metronome-Industries/terraform-provider-metronome/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* add per-resource api permissions to schema description ([52dd72f](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/52dd72feaf18954f910ad92a15bc35fa8cc2bdd1))
* added capability for `dynamicvalidator` to do arbitrary semantic equivalence check ([27496ea](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/27496ea4ee44faadb3e51caf214a9a6e70efa8ac))
* **api:** api update ([2c7a26a](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/2c7a26a12e3e1d0dc9ffc6132e1ce296d030519d))
* **api:** api update ([3abaec4](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/3abaec457971f14bd6065ff7c4bca1a6284cefe7))
* ensure `internal/apiform` encoder can handle "force_encode" serialization tag ([bb6dfdd](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/bb6dfdd573b1b79d81ce127411e0634b2deb3001))
* **internal:** support CustomMarshaler interface for encoding types ([cd590ec](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/cd590ec47896d30046299e6863f812352f487d64))


### Bug Fixes

* **api:** handle mismatched dynamic array types in state and plan during serialization ([3df79a0](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/3df79a04a9ac98327ae1a8a93d0a019f44c4ffd8))
* bugfix for setting JSON keys with special characters ([6fbec92](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/6fbec924fc818532231e639f1ec08918c8a17b70))
* **ci:** in custom setup-go, pass through go-version and cache-dependency-path ([ba14683](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/ba14683e671845dbab2b41238dca6b3d78c32735))
* **client:** correctly encode map patches ([e1b760f](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/e1b760fcfb056f8f2d6731ffe6d8d7777e3c73ee))
* **client:** correctly patch `null` -&gt; zero value ([56313cb](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/56313cbff0d90d4879fde3d72473caf147869cd8))
* correctly detect more ID attributes for data sources ([f6bc12b](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/f6bc12b2b964723ce4df2ba573c17ebbf595fe42))
* dynamic type validators should handle int and floats correctly ([f459e17](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/f459e17f35c30248fb4912f83ec1b876ac8a6d5f))
* encoder crash for nested nils in dynamic types ([7603da4](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/7603da41fa42e5b4387afdc9828d89c9ade751bf))
* ensure dynamic values always yield valid container inner values ([9be83fa](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/9be83fafd1f057b2a5a34d70af1efeda41a3c385))
* fall back to main branch if linking fails in CI ([b6f56ca](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/b6f56ca7214ef3e5544c99b70f814a13f3791b7d))
* fix for failing to drop invalid module replace in link script ([03b7601](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/03b7601f4d051a5db9ac54415b99f07f689ea1ac))
* fix quoting typo ([0d05804](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/0d05804e46bfdae2899d226052781d476d30e399))
* improve linking behavior when developing on a branch not in the Go SDK ([636325d](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/636325d4fab3cd64cc87bde43989a1b6c79a6123))
* improved workflow for developing on branches ([0947723](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/0947723af4281262ecd797820354466dd233391a))
* list style data sources should always have id value populated ([324e551](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/324e55134e735c1736f960c4acbdba9870195be9))
* no longer require an API key when building on production repos ([246e3b9](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/246e3b9d981ed084090591cd714e629633ed520e))
* patch style requests should never send empty json body for objects ([43d7c83](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/43d7c8357a0fd89fc63e849bd289496b87f2e91c))
* populate computed_optional collections from API responses ([f3f117d](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/f3f117d27aa2062ec8324e53b8c63a8a3ca729ba))
* properly handle null nested objects in customfield marshaling ([981ca2e](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/981ca2eca2b30528cdb551327183f220843f5b8f))
* **scripts:** export GOPRIVATE and use `go mod tidy -e` in scripts/generate-docs ([9dfd2d6](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/9dfd2d654b430b8da97302306623805767765d1f))
* spurious update plans for float attributes after import ([4dac98e](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/4dac98e181d3dd4460ebfb3b0c2c18dc3b8edb40))
* **tests:** update hc-install to fix PGP key mismatch ([badb38d](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/badb38d54cd90e52a61fac5af9a843c21ad08892))


### Chores

* add local tmpfile directory ([9323ef8](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/9323ef820936cbd5b664b46e099911698b437cd5))
* bump dependency version ([d7cd825](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/d7cd825aa104c00dc37eee62c9c81c94fbe0c262))
* do not install brew dependencies in ./scripts/bootstrap by default ([ea5b4f2](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/ea5b4f2fd13e8771a8c20fba86e55be03d53076d))
* **docs:** update terraform-plugin-docs to v0.24.0 ([487ab98](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/487ab988c81d18cfdec3b801736ac02db1fbb178))
* ensure `tfplugindocs` always use `/var/tmp` for compilation on linux ([19aede2](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/19aede21c73a82bb59cb6d34fe08bc91b2466cdf))
* ensure tests build as part of lint step ([cc64206](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/cc64206d46d15e84210b0a1e7b26ccdecf31326a))
* improve integrity test error messages ([1e3ee53](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/1e3ee537876e2d30f2ba7fb95dd0092f4cb9e3b7))
* **internal:** add test rule to lint for dynamic attributes that do not have planmodifier ([5d4c9ea](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/5d4c9ea0ddb66fa87d920447daf8598fdd211633))
* **internal:** address linter warnings ([e9efc3b](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/e9efc3bebbfd11e3d3c6cb4e95eb31e02bcaa878))
* **internal:** codegen related update ([731d058](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/731d058fa4f0985482300495d7578315a01afcd8))
* **internal:** codegen related update ([adf055f](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/adf055f694e88d38e5c9cd074e875515a90b2364))
* **internal:** codegen related update ([f9fe00b](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/f9fe00b80d1cdb340f7dcdd2d4689dab3bd3b382))
* **internal:** codegen related update ([3f4b6a7](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/3f4b6a76f0f050c75572b3674f1a8606c323cae2))
* **internal:** codegen related update ([a98b163](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/a98b163a4df819914e09a317100cadad7897bfa6))
* **internal:** codegen related update ([c28c123](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/c28c123aa5dc3bba5e6d92ed6be62f511839218e))
* **internal:** codegen related update ([3b1bcab](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/3b1bcab510ced0ad4d1e1778fc1ae7971847bb1e))
* **internal:** codegen related update ([2a5af4a](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/2a5af4a71ccc38b70eceebbb90b091c5b298487c))
* **internal:** codegen related update ([2840520](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/2840520e21248122e7a53b9dfdf861943e00a171))
* **internal:** codegen related update ([88ba9d5](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/88ba9d556a9821d0e37a895c0090f6c9b87355e8))
* **internal:** codegen related update ([e17f4e0](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/e17f4e068a1a261ce4771bbe92e65560d7e0d39a))
* **internal:** codegen related update ([4eba399](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/4eba399a9e316572a54609d75b10276a5404bd6c))
* **internal:** codegen related update ([5cd7cd7](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/5cd7cd709b3b4bbba21be69dd0025f4189cb855f))
* **internal:** codegen related update ([abd20dd](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/abd20ddc00074fa9fd06406de0d2687345f82fe2))
* **internal:** codegen related update ([af1938f](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/af1938f1d6322ba8e4491d39a5a8d91925c9a0f5))
* **internal:** more robust bootstrap script ([70f3c3f](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/70f3c3f42d229a89bc2a2cb2148214df5f40616c))
* **internal:** refactor the apijson encoder ([19e91ab](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/19e91ab8ca83b940a91597692b2c0d90b8e2e1df))
* **internal:** tweak CI branches ([0dcabc8](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/0dcabc85dabd46ae66ad01603e7b7acb22ef3c23))
* **internal:** update `actions/checkout` version ([9041037](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/90410370a6b5275e621a66640cd85888e5b2e335))
* **internal:** update `interface{}` to `any` ([7b18109](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/7b18109cd42d02fec0f1ac322c70326b7c3df498))
* **internal:** update gitignore ([86db47e](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/86db47ef6ee78ae3efdb1d4c53ad796a445bea94))
* **internal:** update multipart form array serialization ([dd34eca](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/dd34eca8b77a751ff29e8d9587b94704191207c3))
* **internal:** upgrade cloudflare/circl ([5d33889](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/5d33889d1fbee8d0b10e215edc0ce300eeb298db))
* pin go releaser version ([8ffa28d](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/8ffa28d94501af62c96099d93fbd4f705b944f9b))
* **test:** do not count install time for mock server timeout ([bbc6adc](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/bbc6adc48d6a2bfc17e67225fed1db6a579997ce))
* **tests:** bump steady to v0.19.4 ([929212a](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/929212ace7427f2f86705ce89ffc56e9ccde6a95))
* **tests:** bump steady to v0.19.5 ([c20e311](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/c20e3114c6f513085e47ea3ffbe952a4ba399471))
* **tests:** bump steady to v0.19.6 ([baf1e71](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/baf1e713b859a8fbf90dcf02aa6a951ef0c69cd5))
* **tests:** bump steady to v0.19.7 ([e477fc1](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/e477fc111fe66d8d97eb95355aed330aa93ed7c5))
* **tests:** bump steady to v0.20.1 ([8334ae4](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/8334ae43788956f724b7835f4ccc22c9902f0059))
* **tests:** bump steady to v0.20.2 ([939bf39](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/939bf3935c426db32895ee5b9a3cba6b06454197))
* **tests:** bump steady to v0.22.1 ([b7a7708](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/b7a7708ddeafcda56d296386ab02567a5d9d69f8))
* update @stainless-api/prism-cli to v5.15.0 ([2486297](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/2486297fad4f421ba83ef98745cc8967e26e8f01))
* update terraform-plugin-framework to v1.19.0 ([2f5bf7a](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/2f5bf7aca57310946f29639daad2cca06fecc10d))


### Refactors

* **tests:** switch from prism to steady ([77af9a6](https://github.com/Metronome-Industries/terraform-provider-metronome/commit/77af9a69ae4f190521c74ec4d59aa154636346af))

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
