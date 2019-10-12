//
// This file was generated via github.com/skx/implant/
//
// Local edits will be lost.
//
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

//
// EmbeddedResource is the structure which is used to record details of
// each embedded resource in your binary.
//
// The resource contains the (original) filename, relative to the input
// directory `implant` was generated with, along with the original size
// and the compressed/encoded data.
//
type EmbeddedResource struct {
	Filename string
	Contents string
	Length   int
}

//
// RESOURCES is a simple array containing one entry for each embedded
// resource.
//
// It is exposed to callers via the `getResources()` function.
//
var RESOURCES []EmbeddedResource

//
// Populate our resources
//
func init() {

	var tmp EmbeddedResource

	tmp.Filename = "data/archive.tmpl"
	tmp.Contents = "1f8b08000000000004ffac54c16ee33610bdfb2b6689343d74256e0a1428124ac036692f45db45b3973d159434b2885024418e9d1882febd2025cbb29d5b3b1791e2f0e98de6cd131f9efe7afcfaedcbafd051afcb8d880fd0d26c0b8686951b00d1a16ce2024090228de52fda6ee1b3af3bb547c1a7771b88e73d9284ba933e20156c476df673c200988e8cecb1607b85afce7a62505b4368a860afaaa1ae6870af6acc5ee3e62328a348499d855a6a2ceef24f4728adcc0b78d405939ad01b49c8800e0e0b269dd3aa96a4ace13e841fde7acda0f3d8166c18bc345b841bf5116e08ee0bc8c77118540b370ac6711850071cc761c8bf786cd55b7a659ae5a14c836fb90f81412ab9607f3f3f438bd8406b3d0cff117e2e6e18087ba72521b04692e4cad4bc0e21a7de69368eb121fcd81151d9e630fffa46eda1d6328482c563f433208090ffd31f60c70f68bbb5ac7c26dc23fc8ebdfb3e40d484084e9a123278c24a4903b7b2770ff09b478467dbd2abf42878ca115c4eb401ae89675e6d3b5ae8af0ae0b2b23be2acfc1c9f2b909433572f6b527b6473cd5c4e3a8d97a6d5e5b52ea98393dc06cecaaf721bce32de6d883f6b480cc11bb59f6a5a2f5d796baae01e0477f321c94ae3b48e03e54b410d4cf267779fbe5b97cd699ebb1862bd8957bbbb72a9a8bb3b22a648ff34d041e33c59f73f9d21032c72cd93a44e2176fa7853ab405942b937d6e0c38ada1442ab5254e530e4df50fa7114bc3aa3912202ce9d49805b6f77ee0a6acde84992bc603585d0ea1a2a53843d2b17912f0d5f68c5d51fd650378e9c45b669f3a7ec3152962524dd1e912bd96c31a53dda9da19892240b826b755d5e229e6c2239e02904dfe9f2fa4e7a7d963a5c5f5f2928c695108e7af9f1acabe243964127eb17c8b2d327def5a585c0301cfd2ff2d0d10097b3289277ccc8638d86fe7136d0ec4a1722ba2ce872bf2e4770f21357c15793f1cedcacfcadb596d6fee68ebdd3d8122b1f3d4ac206aa03a0ebb047af42be0c20c0297f369adbdabac3c3c9273b2217ee390fd1e0f217ec5dde2abe36bce8100be2d22dc12bdb1cca8de01df5badcfc1b0000ffff9d3d4b045d070000"
	tmp.Length = 1885
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/archive_page.tmpl"
	tmp.Contents = "1f8b08000000000004ff7493cf6fd33014c7effd2b1e96180768bc7242cc8934b671418889eeb21372e297c49a635bf65bbb2acaff8edca45dca8a2f7989bf7e3fbef958bcbbfd75f3f0787f072d75a65888f400236d9333b4ac58008816a54a0180204d068b6fc635701daa566f1096d0f7d94f67a91d06e8fbec11651806c1692f5d403ad62149a85a192252ce9ea95e7ed9a70618b7acec30671b8d5bef0231a89c25b494b3ad56d4e60a37bac2e536bd7c026d35696996b19206f35576794865b47d82802667d210062b0919d0ce63cea4f7465792b4b33cc4f8f1a5330cda8075cefa3ebb0f58eb9761d056e14b16626440a9fb9cfd5eafa1465450bb90e63c28a7927d4fd8792309812949926b5bf12ac68c3a6fd83024f7f8c13e513ab59b0c517a03959131e62c6d6398120208f9b62f76d01ad73856ac0937083fb0f31f22a47f21a297b68025dc62a9a5850bd9f92bf81e1061ed6adaca8082ef3582cbb10380b73d2c836e5a3a7632eb85cbd23d1367c5757ace929c684620926a8cfea323d944ce8a07d9c413c55933c389996909aef4661c621efae2c296d15f09eea74d92a5c1314ee4864290821128b6ba7c3f9f93d304785a62fe92f05f15779682c608de454205da2616ce30dfae0ef5d2eafb206d83904dc7f73c1cd6d9694be39a3fa9c8081064ff1c41ab665fdeb47d98eef3c97470be56c00a2dedab4dbc9e949b27179cc23898e03357cf783ee3ba768ee65cfb23c458132b6e02cae465b903f42d761874cc8e3f0fe0553f51795139bfbb7abd1f2d918f5f398fe936644fd8f9acd67c7e3b125dc78c4756042f9dda150bc15bea4cb1f81b0000ffff501ca54802050000"
	tmp.Length = 1282
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/entry.tmpl"
	tmp.Contents = "1f8b08000000000004ff9c53cd6edb3c10bcfb29f623f0a587d662d253d15002da24bdf4d0a0cea5278312571211fe81dcd83104bf7bc1487664c4b994172dc5c170677628febbfd75f3f0e7fe0e7ab2a65a88fc01235d573274ac5a00881ea5ca0580204d06ab6128ee1cc55df190b7fbbde0e3ff05648c4592d0f43226a4923d51bbfcf2c203301e3969b1641b8ddbe0233168bc237454b2ad56d4970a37bac1e5366f3e81769ab434cbd44883e5557179a032da3d424453326908a393840c6817b0643204a31b49da3b1e53faf86c0d833e625bb26128ee23b6fa79bfd74ee173115362405946c97eaf56d0222a687d84e115395d390c84361849084c49925cbb86372915648361fb7db68a1fbc12b557bbc910a537d0189952c9f231c6891040c8b77db103d6f8ceb36a45b841f889367c48f0ddf84ea4205d054bb8c55a4b0717d2866bf8111161e55bdaca8882a78c115c8e1d00bced611975d7d3b193592f5cd6fe8938abbee5ef8ce404139b5e6f30a3c6ea1d1cc92e71563dc82e9d20ce9a194fcccc4b70a537a3887919aa0b57a7702d78980e49d606c73ac7345682148c81625797ffcf75729ad29c97986fe07c5bb5f1dd3af844e3a4610cffcbc00febac9cc65b8b8ed2dabbf5bf7248a5d613cfbaf5d11e5a98ddfe46d041f7e713ddef888bd8a0a317795392e13d72c1298e160b3ef3fbcc3466896fbda779e2c331ded812ab6e224a4205f50e30f46831ea541cc70af08a9ff27ad1f8b0bb7e7d393d51485f394ff99d148f6843d16a3e7f37397747c6638a04afbdda550bc17bb2a65afc0d0000ffff4f98d5ce09050000"
	tmp.Length = 1289
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/inc/add_comment_form.tmpl"
	tmp.Contents = "1f8b08000000000004ffac54dd6adc3c10bdcf530cf3c177d5ac9394f457368492bbd084b6107a2959e3b5882c19699cc4357ef7a2b51dbcdb1f525a5d2c23f670cecc9923c3d3190653c1e642eb0fbe69c8318ce3d1d39f20eab3e2426be87d17a09c1022abcf8a35a6f2a10159b2f12ec7acdc9a63655c36a3e3a6dc1a04a3732c1310c1c9869e2e0d71ed758e37d79fbf20c8b2a4968fcb5a86489c63c7d5f11b2c566200c2b8b663e0bea51c6ba335b985d368847b693bca7118aeae6f2f3fc1e6d271e83757c6dd8d2342b66efcb75cc12bcf53df733933ff0949157ce47e6259ea99e6ffffdebe7e79fafe703896cad27e9300824321581742151f65432253c53b91b13ec425a42ef60c627a648468be518ee7e78b516905698e1d89c8381c32ad142f1b69ecdf4b2696e74ade928a86ffc19cd6b8bbe78842e96d6ca5cbf10c93cb5f53dce7079146172a1c46673a22f92b03c939d5caeb1e21f88798e3e909ee78737c7582c9657adc4153a9d3cfaf5d4fbb15ac415ab3753906b3ad19f7171b3bd5185e16badce670ddf8c84bff9825ad9f098aec87b4892cbdd1751a443bb9b15388d178070fc65a500466eb7c200da602e97af015704d5019b23a820c04962a0665a5bb7b01aa63d87d45282501a4d681629cb8842a1cdd53485627626d626b654f7a23b276ddcc3090d3e378f43d0000ffffb645fce9bf040000"
	tmp.Length = 1215
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/inc/blog_post.tmpl"
	tmp.Contents = "1f8b08000000000004ff8493516f9b301485dff915575eb6a706b23e1263a94ba2bd444db546da63e5c20d58750cc3776d23effef7c924a4eda66e3c61f9f89c7b3e83accca34a00647309a5d5de17e2a7abb0b7c6612594d4d0f4b82b4408ebcdf7d53748d7c63d300bf074b05808c2679a5658b6bd26d3badcb50ee750b6b6edf37babcb87b95021a45b43169965a695cc9a4b9500c80e3c1d2c1682f099a6da9adae5bda91b9ac393a9a8c93fcf661f850ae1cb7af3f5ee6673bbbd5b5e6d57902e3521b3cc3a9500c8ca3c4269759c1c1df58769d93a42474285902e8eefccb162f766686d5f029f4c45635e0200207da7dd08e464ab09ebb63f88a302428089455753037901161da45b5d7b608e5b6607f8e32c98c130c1f05cb740baf6434c54a2f5f86a3b7ae41042af5d8d30311730a1183098338760763031cc171002ba8af97c4559b4cd4258dd2eae6e56903247042f2b9969753a33860f06c34266b1f1b1da9bf27b243df5d809f5eb5dcd0950bbdfa323ff0f408ba364a949ff8fd3750ba3e130e0dfa8cec5fff8363f8ce762fde15e5ec5329f6d238fd1fa1d1059a712990dbf88ecd42777efbbb9cc3a95fc0e0000ffff4f04f8f337030000"
	tmp.Length = 823
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/inc/comments_on_blog_post.tmpl"
	tmp.Contents = "1f8b08000000000004ff6c90416b1b311085effb2b06d16b778b7b2b5a411afb5048ead21a7a0cf26a6c896a252a4d1ccc30ffbdec6eb324e0b969e67def89c70c1f22a63379f8d243c404ed7d1e474cb4b564410498219c00ffaeba4f20c20c182b8248a3fd0682ebd5b0605541b223be799bff86157202f2a102262a57ddf98d6998313991a6612e369df15dfce4eec20586686b5d1d9569006eec3d5a8765be82267b8c38094153319a1cbc0447be579f378be2757418cf6023f52a0c39a93732f018ce9e66046a197ac5dc7e1b7212510bf15cb17c9c31a33b7273dc349a9c616eef9ec9e7325778bf7f7cdc7d3f3cfdd8ff3a3c6def0e3b68b79650441f4bb772f3680bbee0694a7bd8ffdefd84f621a43f5366c1d8ab944f39c6fca20cbfbfebceae46eb6f744765deea6e6d44772e5ca6e58d0e8fd95d977e98dbafd95d45a6b21742772e5c4cc38cc98934ff020000ffff0624259d3b020000"
	tmp.Length = 571
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/inc/css.tmpl"
	tmp.Contents = "1f8b08000000000004ff7453db6e9d30107cf757ac54556a1fa087b4912a1345fd15c32e60c578a9f1b915f1ef952f3490a46f30ec8c87d9f1d3ecef869e4539904272b00800be90eb0c5f250c1a916c2d001ad5bef48ecf168b960d3b099fbaaaabbaaa16ab283359457a67587909863a1f986d1a6f8c6a5fc2bba79b2f94d1bd95d092f5e4023a29446d7b09d5c3740b401c436ad929afd94ab06c297ce8d8fa62d67f4842f533cd1a6da91848f78397f0f098c0861d922b9c427d9e25fc08e8ce6a69b8e7e4f75570e346e89a051b3658bfa7ce93b2eff8a7700a1cf996dda8cc41410e21e3c8fe2058440c3f7a0c6e6f5db55e5fe83f7464ecc935e61cd3ca22d7417bda59285c082b2ae47545204cfc1a09b582b97544169445f832aa5b71d5e807098fa7d374fb0a8b80f2b074d874b63501a09e27a3ee121ac369f590b69a979f0bb28a4d6a67eaaddc2a56e111965d4f42d47021e775abcca6e979aa61155e3586608164ba3a9d3ed75b1fe649b5a9685120b7a46563d43493849926e594a720138dff3eb30f5a7932b89670aa372719788c6aabc09dc9b79f3c2cc766c466fd13f23cc9d05d5845d931fb701b21b7fa7b94cf16c26035dd6066a3315fac400ad660d9671c903a6f263dafa28c9b5e321a5fea3d2723ab28d3ed243c4826309e77b6482edc3d84e583a41bf69e4709af56676d2e94b82d8f23d96037fdd47eac7774af6154aed77653a94232c153e2a5ee2d99bc0ded8e4a1aaf8486f1febe3eab78fa36fbbba167f1370000ffffa77233f909050000"
	tmp.Length = 1289
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/inc/recent_posts.tmpl"
	tmp.Contents = "1f8b08000000000004ff84ccb14a04311080e13e4f11d29b14d7e9dc8078e90eef58039647564737386461932d64987717a2bddddffc1f2c079ce88d6ab7d7b5f5066139a0819d6debdf4c47c7a5f5bbd1f775adf4e0d0886cb97e92f5bfe3f8540d70419851648a4ff139ddae9797743b3da668fd2977528530ff470f23db65a38fa313395f5ee364fdb9d42f5587223e95ce83ca08810b1a083bffa508d57755036167343f010000fffffee417f3da000000"
	tmp.Length = 218
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/inc/rss.tmpl"
	tmp.Contents = "1f8b08000000000004ff04c0c7aeabcc0100e07d9ee2e86c1d858e21f9ef95064cc7f4ea1d0c43b1e9dd3c7dbe7fb29f7a46e59f5face90b74fe675e96df9fb5595bf4e7d7f3fd1f19a1e2f7ef3f4d57fd2c33fcf35b646bf6dfa6cb2a848d7df5bf3c5b104bffbb8904db3b7043a9060000b0fcb096c20a0020540000908a2005003caae2a2680080bce1ade4469ed6a4a41fc4077c9dbdda5345579cf83d8a37df1074914981bb1ee3e9d966e81e37486b8f0a8888d67c591eba0f5c29b5d087f9adcf91deb59e95bd5752d65be948d7a5daf00cdda2618655191436fd1845cdaa850691b9fabcf67741bdc85de0b0c7ec20a5727d0d235bbcb46e987ec730aa953dcbeb127ed929612d6fd0cd1df6cb73d0ae2d8aa64d9bf57676d39d780ef9ed9e3b6847f62e1e9c7a9f254ff6d299a4ee6489484d3620d2dc3357c9713f171b57c81e3d288c46df6f79ef6ea5fc4056f790c12b7f0fd585cff992ac37c7e645b1e4964f27e74397e69d6c2b505b8ef1e85b99856c226e28d84626601f5c7ba3adb10ef88eae589c9689f9b893b73b659dc26ac7e81b3db8d74b7f8ae58042d51163d936445c63c2b04eeadc40f7041fc3c974fb3a1105289521d4e0a80367b9aa871189a4371a31cf5442cffb4362899f479926d6e4ac85a7682f4ce8bd58245d1f87f84d734866185d1949932e0f8997888ea557fc1a881d3754d706f1aaad8ca743b6d2c02624d946bb3d12fe2e71b1853f3591c25a12de38ac66528f0eb33794a917e9969cbcbcc93767301d346eddaa368fc94fab5eb08cfe1ef9d3f515d010f3c93d8c25b556a147736f864e8c1d497d44b94b0df6bb2b3f0d7a7c28c5d7db5341f7925a6df4d4a23fe59d2cbba647b342b03de3b03a42852b11f170c3b69ce26c3a8d1eccb629bdb4a2552c79173c1e109c50c249de36f041734df77d56b37819b29fe879fb31436265f93d9c287eae660f933a6e3c38e67b70c4576a6f380e0418a0a7dd1113078ae7f49872ef438f7e52bd2fb1297a7a6127a617e6fe98881b6eaed0df094ead82b01a55cecb8348cffce93999eee826c7c0856726e1d42c8ea7b7e03332d4acf4f1617483a182afef2a3d43d452c29376a0c37c2591df97f9d910328c56c83fd945713a061cacf8d08cef64d9a16b8a1094fef6b4eec144f7d20016dce0888fce5ff46376c50f3ccaf6b2d9f8fe80cde426d45de728227588bc55757db6c34ed343b036ecabede179966929dbca2289f8eb943827cc63cccfa8702711fdca4fac1cac434bbb651ba746deb913c572a460187296f972bb169126bb1a3263d3114dae51cdd75b0c8348e566df4825c7c72589d6f1f47029220afb93fc246dab6f13f9c0682e69125d2c426570947677b6d245af4f65dcaaf1c86f2d74d6407ed2793c7daf7a8ba1d5b428e43ac5f71528651aa3415ae75ea78b8a5a745ecc6283fe88f84d63189be8deacbf191577c0adc7e55e508990defd4c763eb659c0ece8cdc26a2adba08dac3a8f080bcc87224b4c6c3c715cd3f86c2a2004cf39d779d1427b7fc9cc700a3dcf57b1101a809e7bfdd9d5b9adc8dbf6498b76efe62a914b9a38bcc5ebcda46fc2e95d5beb4e672c2e3d718d89bb9aedf11e468a83bb7cdec4b4398671106fa10074f5fe5ce9e8280940e7e7988e7c1c4fe776dcb57b90e7ac5c30f2f1e28ec93eb84c4d2394dd2afdb4c0ad783ef96418bfc9c57bd05806853361295c409cdbeeaa4a5161b163a3f125e44633caa8653b6d702b8a3a29044e822181443508450e6b233c6a6a4b4cc6384cf7cb0f2ff9ceda97fb06d369d173933b86b23144c51297b86ff4c705e4c678ad7a1507aa131b417595b0e3ac303338181fb3358024ccb3dc18ca8eaf84fdc97c5eb7a51d66aed4ccaabe015b2cf70a4224b5435794c74978b6891b1bef807bfbcaac790a5f4998cae610051356f1a0df6ca8ce727e5d55cfcba294f9c1f0ccdfe5adae3fa95fa985c0a983b4e2c571559dac0e095f0f6aa467ca5270b3f11a4ad19eb2633b13db80ea28152725187cc9a96c90d82286ddb0aff90eb7fb87cbe0caeb579da0706d76833d2859587a46fb70589e70856d0fd1a3402151ab72f99ac5790df494b18dd771893b5a8e9ea21ccadda660e1c5ad58eb8410d3040e33606522eb8ea079c7f62d4ab4c69cc86cff7eb1e8bdccaf33799a9659756b5e7dccd7d75d7abd7451c88b0c6f62b814eabdbbf8f23e7c5689373335d23747616b61571947b9a25b3c34468d01c4734e67c45d4b37000000a4560e3efee676a2f8fb93b5eb9f5fcff77f64848adf1fecef3f58f6f75fff0f0000ffff0efeeaa54e080000"
	tmp.Length = 2126
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/index.rss"
	tmp.Contents = "1f8b08000000000004ff74514d6f1a3110bdfb57581ba9373c40a4aa6c8da3368054355222a8d4b3b39e2556bc36b54d96cadaff5e79592829ea75e67dcc7bc3ef0e8da16fe88376765e4cd8b8b813847b5597ebc58ad043636c28bdaae7c54b8cbb12a06d5bd6de32e7b73099cd66309ec2743af2aa1e85df36cac3c8869b820c4455cd4fbcddde9b9ea52a40830dda1860c226509c5c6a27ffdaf4c6ac720de4318c2f8195b3116dbc96f6214b8ea1716a6f30c0003c5bfc9f51104178f522ad4543737af9ecf6715ea4c49e3cd6fad07585203cea68506c22be21fd8ecdeec3cdedec73a05f8ddb72382e0937dabe8a0b22877e42b8c25079bd8bda59b1c0672d2d9556d19547a41b57c7567ae4a02e5084eb884d1084f62fd9e02f4152f2d26e91b2a58d5e63e83a428f6ba3fbd33d06b7f715ce8b941e1e7f2ed7943d68fbda750564365a95291c72cc5e91c3e0c261a820e3ae5cfa5b7a877339efe5cf05a5b4dcdc7f795a52f62317d67557e5bc239e0b1afe55a2ad9c4275a1737f7c7956fa1744b8aa4a25238a94be6d1e3f7d1c4f285bc8d8db9e5643c69ceb989f835775b95eac04f9130000ffffc1b1f9e704030000"
	tmp.Length = 772
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/index.tmpl"
	tmp.Contents = "1f8b08000000000004ff7454c16edb3a10bcfb2bf61178e9a1b598f45434948036492f3d34a873e9a958892b89084512e4c68e61f8df0b46b223372e2f5c7ac7b3c3e140eabfdb1f370fbfeeefa0e7c1560b9537b0e8ba529013d50240f5843a17008a0d5baa564c6b82ef34847709be5adf293936169041033142d3634cc4a578e276f9e98508606c391ca8146b439be0230b68bc63725c8a8dd1dc979ad6a6a1e5261f3e8071860dda656ad05279555c1ea8ac718f10c996022d5374c82480b7814a812158d3201bef644ce9fdf36005f491da52ec76c57da4d63ceff7c6697a2e624a02385fab143f572b688934b43ec2ee15398ddced9886609109844646695c239b940a1e8215fb7df64a1ecc52b5d7dbc9106dd6d0584ca914b94d71220450f856973860adefbc3863760ae82a58c22dd5061d5ce010aee15b2482956f798391944c19a3248e0a00de6a5846d3f57c5432d322b1f64f2c45f525ef3392134c6c7ab3a68c1aab7fe018bb2445f5805d3a419c35339e98999792daacc74bcccb505db83a856b25c3d464ac2d8d75ce69ac146b180325ae2eff9fdf53f214e7bcd4fc00b0db45741d4171e7381a4a2faf7a586735d7d677bf834f3cc6008abffe424ecf7e7933fca0f1e389c62ce44cd82235e4f865da94ba9371737225398e762839f3e68c73b374b6def33c9de118456a59543791904943bd050a3d0d144d2a8e4f00f08a9fb275d1f8b0bd7e4d79cf1cd2672953fe80148f3484a235729ef19c9123e3f1c595acbdde560b257b1e6cb5f8130000ffffc02a5f07b6040000"
	tmp.Length = 1206
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/tag_page.tmpl"
	tmp.Contents = "1f8b08000000000004ff7493bd6edb3010c7773fc59540d3a1b518772a1a4a409ab84b8706b5974e05259e2422144990173b86e0772f68c98e5cbbb7e824fe795ffa9d78f7f8f361fdfb69092d75a69889f400236d9333b4ac98018816a54a0e80204d068ba5a5a031c25a360d2ae8fbe5eae1fe6909d95a36fbbde0836a06e9468724a16a658848397ba17afee510156038b2b2c39c6d346ebd0bc4a07296d052ceb65a519b2bdce80ae7dbf4f209b4d5a4a599c74a1acc17d9ed3194d1f619029a9c494318ac2464403b8f3993de1b5d49d2cef210e3c7d7ce306803d639ebfbec2960ad5ff77b6d15be66214606947accd9afd50a6a4405b50bd0bf29c7947d4fd8792309812949926b5bf12ac68c3a6fd87e9f06c78f9313a553bb71204a6fa03232c69ca5630c634000212feb6247ad718d63c58a7083f0033bff21c237e31a11bdb405cce1114b2d2ddcc8cedfc1f780082b57d35606143c268de072a800e0b28679d04d4ba74a26b57059ba17e2acb84fcf4990334da85abdc1a41abcffe8483691b3622d9b78a6b83acc7036cc64822bbd199a98babeb8b165f47782fbf190646970f013b4a110a460008a2d6edf4ffbe434b29d4c4c5f12f98b13eb34b02e2aa7b0b820fef055f076714c9aacef83b40d4236eecb018aa35d6db934aef9e35da48122c8feb982564dbe5cd47e6cf1f3598b703d57c00a2d1db28dd09ea59b06179cc2d098e093d15e19fc04eeda399ac2ed4f24634dac7808280915943b40df628741c7ecf40701def4239a3795f3bbbbb72569897cfcca794c2b913d63e7b35af3e98a24c44e114fc0085e3ab52b6682b7d49962f6370000ffff347ad4c102050000"
	tmp.Length = 1282
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/tags.tmpl"
	tmp.Contents = "1f8b08000000000004ff74534d6fdb3810bdfb57cc12d8ec61d762bca722a104a44e7ae9a141ed1e7a2a68692411a148821c7f45d07f2f6849b69ca4bc68e8791cbf79f346fcf5f86db9fef9fc0435353a9b89f8012d4d9532342c9b01881a65110300418a34666b59c152db6d2178ffc30c62b2419290d7d207a4946da99c7f3a1500e853463698b29dc2bdb39e18e4d6101a4ad95e1554a705ee548ef37dbcfc07ca285252cf432e35a68be4762ca59579018f3a6552137a230919d0d161caa4735ae5929435dc87f0efa1d10c6a8f65cada3679f658aa43d72953e021f12130a0d84ecabeaf56502216505a0fed0539fc65db12364e4b42608524c995c9791e42428dd3aceba2467c14496c6c711c0429d40e722d4348594ca31f0a0208f99e171bb1da5696652bc21dc2576cdc3f013e6b5b89e0a4c9600e8fb851d2c08d6cdc3d7cf188b0b225eda547c143c4082e7b0600ef39ccbdaa6a3a339970e17263b7c459f610bf932257189fd76a8711d5476f7143c33227b54336b4c9495681b3e89c70f5e0436dfd95b6f1085ea85ddfd33474d98dd904772fb81b9224371afb38dad567820ae8fdc516b77f4fdbe634b83a1e31b9887a11898e16af17633d685b2f4d8590ac65154e833f9dd36420d05163ca4a6b681ed42bde452bad57ea15bbce1d58769e792f46db3ead960fcf4fa7625d7736e38f80457cb9b45b435d07a41a0c2c7b038f228a7edaa7e583480e4d712105ef1a1c75f8ff4a87f8f0037f7bccd1d02f67030d4687e40fc50527dfb3107ca2ff07d3992c44692d4d17c28dc6d15812cb961e2561019b23a0abb141af42721e33c0053fd8f926b7ee787f59ac9ac8853bce435ca3e4051b97948a4fd7aa977034ce6830c137b6386633c16b6a7436fb1d0000ffff870d5fd721050000"
	tmp.Length = 1313
	RESOURCES = append(RESOURCES, tmp)

}

//
// Return the contents of a resource.
//
func getResource(path string) ([]byte, error) {
	for _, entry := range RESOURCES {
		//
		// We found the file contents.
		//
		if entry.Filename == path {
			var raw bytes.Buffer
			var err error

			// Decode the data.
			in, err := hex.DecodeString(entry.Contents)
			if err != nil {
				return nil, err
			}

			// Gunzip the data to the client
			gr, err := gzip.NewReader(bytes.NewBuffer(in))
			if err != nil {
				return nil, err
			}
			defer gr.Close()
			data, err := ioutil.ReadAll(gr)
			if err != nil {
				return nil, err
			}
			_, err = raw.Write(data)
			if err != nil {
				return nil, err
			}

			// Return it.
			return raw.Bytes(), nil
		}
	}
	return nil, fmt.Errorf("failed to find resource '%s'", path)
}

//
// Return the available resources.
//
func getResources() []EmbeddedResource {
	return RESOURCES
}
