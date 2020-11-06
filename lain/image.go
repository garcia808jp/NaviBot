// NaviBot: Discord bot for digital assistance
// LainBot commands

package lain

import (
	// Standard packages
	"math/rand"
	"time"
)

// Register the command for the CommandList
func init() {
	imgDoc := Doc{
		name:        "image - WIP",
		synopsis:    "image",
		description: "WIP",
		example:     "WIP",
		origin:      "built-in, lain",
		Exec:        image,
	}

	CommandList["image"] = imgDoc
}

// image command
// returns a string containing a random URL in the imageSlice
func image([]string) (msgOut string) {
	// Seed the rand package using current time in Unix format
	rand.Seed(time.Now().UnixNano())
	// Choose a rondom integer using the length of imageSlice
	randEntry := rand.Intn(len(imageSlice))
	// Prevent an error if the integer is out of bounds
	if randEntry == len(imageSlice) {
		randEntry = randEntry - 1
	}

	// Create a string using a random entry in imageSlice
	msgOut = imageSlice[randEntry]
	return
}

// image slice
// contains a list of images with Lain
var imageSlice = [...]string{
	"http://33.media.tumblr.com/f3307a1bb2de561902c3f4aa94819b90/tumblr_nttc6uSVuq1u4zgnro1_400.gif",
	"http://41.media.tumblr.com/58ab8b9e534390ab996d0fa15ce0cede/tumblr_notf143Nit1qftmhbo1_500.png",
	"http://animeonly.org/albums/VISINAUJI/Bishi/Volume-11-[S]/Serial_Experiment_Lain_14.jpg",
	"http://animerulezzz.org/FanArt/Serial%20Experiments%20Lain/1/Real/lain2.jpg",
	"http://brendysbookreport.com/wp-content/uploads/2018/06/a1.jpg",
	"http://darkcloudstudio.com/austinifiedanime/wp-content/uploads/2009/11/LAIN26-300x231.jpg",
	"http://dementedrabbits.net/images/medialogs/lain/lain02.jpg",
	"http://fixitinpost.org/wp-content/uploads/2013/03/Serial-Experiments-Lain-Restore-04.jpg",
	"http://i.imgur.com/0fuI3bH.jpg",
	"http://i.imgur.com/4oqKWcM.jpg",
	"http://i.imgur.com/91eok.jpg",
	"http://image.itmedia.co.jp/nl/articles/1807/21/s180721_lain-20th_1.jpg",
	"http://images6.fanpop.com/image/photos/37100000/Lain-Iwakura-Serial-Experiments-Lain-psychological-anime-manga-37187691-1520-1080.jpg",
	"http://images6.fanpop.com/image/photos/37100000/Lain-Iwakura-Serial-Experiments-Lain-psychological-anime-manga-37187691-1520-1080.jpg",
	"http://images6.fanpop.com/image/photos/37500000/Lain-Iwakura-Serial-Experiments-Lain-comic-freak-37540192-411-500.jpg",
	"http://images6.fanpop.com/image/photos/38100000/Lain-Iwakura-Arisu-Mizuki-Juri-Kato-and-Reika-Yamamoto-Serial-Experiments-Lain-kittyluv57-38187394-600-450.jpg",
	"http://images6.fanpop.com/image/photos/38200000/Lain-Iwakura-Serial-Experiments-Lain-kittyluv57-38269174-423-650.jpg",
	"http://images6.fanpop.com/image/photos/38200000/Lain-Iwakura-Wallpaper-kittyluv57-38222200-1920-1080.jpg",
	"http://images6.fanpop.com/image/photos/38500000/Lain-Iwakura-Serial-Experiments-Lain-godofthewired-38580735-387-750.png",
	"http://images6.fanpop.com/image/photos/38500000/Lain-Iwakura-Serial-Experiments-Lain-godofthewired-38580737-500-582.jpg",
	"http://images6.fanpop.com/image/photos/38500000/Lain-Iwakura-Serial-Experiments-Lain-godofthewired-38580739-347-674.jpg",
	"http://images6.fanpop.com/image/photos/39700000/Lain-Iwakura-SEL-rainbow-unicorn-39798201-1109-1281.jpg",
	"http://potpourricomics.it/images/Autori/DeebleNF/Recensioni/Anime/Serial-Experiments-Lain/Lain-Disegni/Lain-1.jpg",
	"http://rubberslug.s3.amazonaws.com/user/3/70f6c9ae3244334822e2fd614703412/ztambtxapy_o.jpg",
	"http://static.minitokyo.net/downloads/15/43/334665.jpg",
	"http://static.minitokyo.net/downloads/42/29/193992.jpg",
	"http://www.animetion.co.uk/Lain/side7.jpg",
	"http://www.japanator.com/ul/29201-1374749821411.jpg",
	"http://www.otakuusamagazine.com/wp-content/uploads/2018/09/lain00.jpg",
	"http://www.vector-ride.com/silly_moogle/TheBigO/Lain/lWatchDadSetNAVI.jpg",
	"https://2static.fjcdn.com/pictures/The+truth+about+serial+experiments+lain+subhuman+_6c62fe_5911108.jpg",
	"https://66.media.tumblr.com/0108cf7e72a70a4adc811a17310a1bef/tumblr_oocj3txIpg1rjtw15o4_250.png",
	"https://66.media.tumblr.com/2792fea59564b6a670549322cc5a8e24/tumblr_p73s6vg9ZA1x9jrg4o2_500.png",
	"https://66.media.tumblr.com/314c1df33c2833397e784288b6a4a280/tumblr_inline_p1ayzt11J21t07u0c_500.png",
	"https://66.media.tumblr.com/68b4d09e30d6eddbde8bc4b85e4e32a1/tumblr_oqv31798n71wqfs22o1_1280.png",
	"https://66.media.tumblr.com/68b4d09e30d6eddbde8bc4b85e4e32a1/tumblr_oqv31798n71wqfs22o1_1280.png",
	"https://66.media.tumblr.com/8a3c407cfb3dfb50fe7255b776a4389d/tumblr_pcn60d1Rh91wjje03o4_r1_500.png",
	"https://66.media.tumblr.com/adbf256d62e42e490d5f3d9707ad69dc/tumblr_p6e3zpsOx61x9jrg4o1_500.png",
	"https://66.media.tumblr.com/bb66b48fdf6a9a41a3aa753c806d4c80/tumblr_pch9cciYKQ1v41zjoo1_400.jpg",
	"https://66.media.tumblr.com/c06286f1360f2550f86217ca4df6ada3/tumblr_oo80im8XKk1vtgpj4o1_400.jpg",
	"https://66.media.tumblr.com/c1b40c07a197eaf747cf06fb3d7fbc78/tumblr_ove6fgVhnV1s73tqio1_400.jpg",
	"https://66.media.tumblr.com/ca94da5c9812e27674bc394c78cf3f92/tumblr_otlbqrT4uq1wqfs22o1_1280.png",
	"https://66.media.tumblr.com/cd549a28872f42b52b429bd46f06b9d4/tumblr_o4ezevquAW1s73tqio1_400.jpg",
	"https://66.media.tumblr.com/d97d01c05943ec198d2724c33e0ca0de/tumblr_inline_p344h9u8Jy1t07u0c_1280.jpg",
	"https://66.media.tumblr.com/f99a00f6c91e0721d84582cc0d60a0b2/tumblr_oocj3txIpg1rjtw15o1_500.png",
	"https://78.media.tumblr.com/eb2c63ac7fb01217539db4f23d676fef/tumblr_osfo1cMJbB1uah3u8o1_1280.png",
	"https://a.disquscdn.com/get?url=http%3A%2F%2Fs1.zerochan.net%2FIwakura.Lain.600.687994.jpg&key=5edZc1IqRe3Z2seRn4Cziw",
	"https://arisuchan.jp/%CE%BB/src/1529427554416.png",
	"https://arisuchan.jp/art/src/1525183091131.jpg",
	"https://beneaththetangles.files.wordpress.com/2010/11/lain2.jpg?w=900",
	"https://c.wallhere.com/photos/15/46/Serial_Experiments_Lain_Lain_Iwakura_anime_girls_anime-229867.jpg!d",
	"https://c1.staticflickr.com/5/4071/4546116303_b8fcd3ff6e_b.jpg",
	"https://c2.staticflickr.com/4/3539/3536914421_7247bdcd4f_z.jpg?zz=1",
	"https://cdnb.artstation.com/p/assets/images/images/005/789/449/large/chrotion-art-lain-iwakura-cosplay-of-neo.jpg?1493786291",
	"https://cs4.pikabu.ru/post_img/2016/05/14/9/1463234747199568052.jpg",
	"https://cs4.pikabu.ru/post_img/2016/06/17/7/1466160572148393809.jpg",
	"https://cs4.pikabu.ru/post_img/2016/07/06/6/14677967391786776.jpg",
	"https://cs7.pikabu.ru/post_img/2018/08/16/6/1534407361174380135.jpg",
	"https://cs7.pikabu.ru/post_img/2019/01/12/1/1547250377171741873.png",
	"https://cs8.pikabu.ru/post_img/big/2017/05/13/10/149469356816223570.jpg",
	"https://danbooru.donmai.us/data/__iwakura_lain_serial_experiments_lain_drawn_by_unprince__d4882ae18c908049ac4e79b3949fa21c.png",
	"https://data.whicdn.com/images/252355722/large.jpg",
	"https://data.whicdn.com/images/308862999/large.jpg",
	"https://data.whicdn.com/images/7353896/original.jpg",
	"https://deninet.com/sites/default/files/images/lain-listening-lg.jpg",
	"https://dic.nicovideo.jp/oekaki/859218.png",
	"https://farm9.staticflickr.com/8323/8095518526_00a4243a4c_b.jpg",
	"https://i.4pcdn.org/x/1494090508726.jpg",
	"https://i.imgur.com/GF7NC4i.jpg",
	"https://i.imgur.com/aOUg9o3.png",
	"https://i.imgur.com/eN7vUz3.png",
	"https://i.imgur.com/qtvrFXI.jpg",
	"https://i.kym-cdn.com/photos/images/original/001/103/057/2c4.jpg",
	"https://i.paigeeworld.com/user-media/1529452800000/511845483bfdeefd1601e157_5b2a1b5eb2bf487d10118e2e_320.jpg",
	"https://i.pinimg.com/236x/03/1e/ac/031eac39a45e13081bb1df7b4de36712--serial-experiments-lain.jpg",
	"https://i.pinimg.com/236x/03/1e/ac/031eac39a45e13081bb1df7b4de36712--serial-experiments-lain.jpg?b=t",
	"https://i.pinimg.com/236x/04/79/c8/0479c8427d6b6a50ccefd8fb7c7fd172.jpg",
	"https://i.pinimg.com/236x/20/66/9b/20669b285f70c4dba3b76118d0184bff--serial-experiments-lain-manga-anime.jpg",
	"https://i.pinimg.com/236x/2e/91/23/2e91237c18ad9822302fc4215a1b4881--serial-experiments-lain-games.jpg",
	"https://i.pinimg.com/236x/75/39/b0/7539b0418e5c187f75eefd1413cc321a--serial-experiments-lain-manga.jpg",
	"https://i.pinimg.com/236x/b9/4b/5b/b94b5be929c8a46ef6fd91f609620b09--weheartit-logs.jpg",
	"https://i.pinimg.com/236x/d7/cc/4f/d7cc4fab8e62a1c247aed7ff92e5500e--serial-experiments-lain-fabric-walls.jpg",
	"https://i.pinimg.com/736x/24/fa/72/24fa72f6573024a94dd43f9f20103892--serial-experiments-lain-white-dress.jpg",
	"https://i.pinimg.com/736x/66/b0/7e/66b07e3adb28660dddd53358f4731ea2.jpg",
	"https://i.pinimg.com/736x/6a/b2/c0/6ab2c0ed376d2a1cc88def88bfe8375c.jpg",
	"https://i.pinimg.com/736x/6f/c0/2d/6fc02daa9e3ce0c72e7e58a4248671d0--serial-experiments-lain-anime-pictures.jpg",
	"https://i.pinimg.com/originals/0c/d8/1d/0cd81d2f357f91244ec87a1d6466107e.jpg",
	"https://i.pinimg.com/originals/2d/bb/cf/2dbbcfb33c21fb240f1715cea11e1f11.jpg",
	"https://i.pinimg.com/originals/51/1c/f4/511cf45997bd698e4f30bfffd7c40bf3.jpg",
	"https://i.pinimg.com/originals/66/92/03/669203fb477f7dd466c25cbecae7d012.jpg",
	"https://i.pinimg.com/originals/68/65/b1/6865b15ca585c3d8f4901d5475c0a342.png",
	"https://i.pinimg.com/originals/78/d3/42/78d34226d2d5bf48a7b71d242b4e9270.jpg",
	"https://i.pinimg.com/originals/92/b7/e0/92b7e038375453b4e541b126b380bb5c.png",
	"https://i.pinimg.com/originals/9e/b9/5c/9eb95c6f4d92335055e26e502ab5192f.jpg",
	"https://i.pinimg.com/originals/a5/ec/31/a5ec310608f5bf394631a04c03aa7cd7.jpg",
	"https://i.pinimg.com/originals/ad/c5/33/adc533dcf3063c1e21702d8f1eb91284.jpg",
	"https://i.pinimg.com/originals/c0/cc/50/c0cc50a78e28b697ec3a66846fcaf520.jpg",
	"https://i.pinimg.com/originals/f6/fc/26/f6fc26e56f252d66197dd8f151c76a80.jpg",
	"https://i.redd.it/dpgtgivinto01.jpg",
	"https://i.redd.it/n5r8cum1mb911.jpg",
	"https://i.warosu.org/data/g/img/0595/15/1490129629325.jpg",
	"https://i.warosu.org/data/vr/img/0051/33/1540878608890.png",
	"https://i.ytimg.com/vi/NXSCGbSXwJQ/maxresdefault.jpg",
	"https://i.ytimg.com/vi/TkXzOlytPME/0.jpg",
	"https://i.ytimg.com/vi/qVMK5WnJ-C0/hqdefault.jpg",
	"https://i.ytimg.com/vi/tffVlvSQOjA/maxresdefault.jpg",
	"https://i1.wp.com/beneaththetangles.com/wp-content/uploads/2018/10/serial-experiments-lain-episode-11h.png?zoom=2.625&resize=371%2C276&ssl=1https://i1.wp.com/beneaththetangles.com/wp-content/uploads/2018/10/serial-experiments-lain-episode-11h.png?zoom=2.625&resize=371%2C276&ssl=1",
	"https://ih0.redbubble.net/image.628018512.8727/flat550x550075f.u5.jpg",
	"https://images-na.ssl-images-amazon.com/images/I/91x56KjbejL.jpg",
	"https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/intermediary/f/1274b4e5-aa00-45ca-a9db-6a90e7073a8c/d4n5m20-06936f97-032d-47f0-a163-d6532648f410.jpg/v1/fill/w_375h_250q_70strp/lain_by_justv23_d4n5m20-250t.jpg",
	"https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/intermediary/f/b841d1a8-36d6-4d82-bb7c-31162693fafa/dck64zn-11a8092a-7ee5-416c-8398-f5a2f0d15697.png/v1/fill/w_800h_1253q_80strp/lain_iwakura_by_peppermint_tea39_dck64zn-fullview.jpg",
	"https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/intermediary/f/e27f75f2-699b-4e21-806d-83a83389356e/d7g8oyb-5ab8ac4e-783b-4883-a264-052c7c9ecbea.png/v1/fill/w_471h_350q_70strp/iwakura_lain___final_by_hackatt_d7g8oyb-350t.jpg",
	"https://images.8tracks.com/cover/i/002/863/598/434e8e4c1145-7107.jpg?rect=00500500&q=98&fm=jpg&fit=max",
	"https://images.8tracks.com/cover/i/010/115/151/lain-4378.png?rect=00904904&q=98&fm=jpg&fit=max",
	"https://img.fireden.net/a/image/1446/25/1446250794874.jpg",
	"https://img.fireden.net/a/image/1454/55/1454557826541.png",
	"https://img.fireden.net/a/image/1459/25/1459259158262.jpg",
	"https://img.fireden.net/a/image/1507/05/1507050296146.png",
	"https://img.fireden.net/a/image/1512/47/1512478457044.jpg",
	"https://img.fireden.net/a/image/1514/73/1514734810033.jpg",
	"https://img.fireden.net/a/image/1519/45/1519455054036.png",
	"https://img.fireden.net/v/image/1499/92/1499927137904.png",
	"https://img.fireden.net/v/image/1518/07/1518074308001.jpg",
	"https://img2.goodfon.com/wallpaper/nbig/4/6c/art-serial-experiments-lain-6375.jpg",
	"https://imgix.ranker.com/user_node_img/50082/1001629191/original/must-not-confuse-it-with-the-real-world-photo-u1?w=650&q=50&fm=jpg&fit=crop&crop=faces",
	"https://imgur.com/061iIJk",
	"https://imgur.com/0NRev1a",
	"https://imgur.com/1a8upI7",
	"https://imgur.com/1tf1TvT",
	"https://imgur.com/23XazJ4",
	"https://imgur.com/2CXVAix",
	"https://imgur.com/2lF0POS",
	"https://imgur.com/3zlHMq7",
	"https://imgur.com/8qfhk8e",
	"https://imgur.com/9IMvVjg",
	"https://imgur.com/AYQARfC",
	"https://imgur.com/AfnSUeK",
	"https://imgur.com/D09hKDb",
	"https://imgur.com/DDlDvHl",
	"https://imgur.com/DECuVWQ",
	"https://imgur.com/DLpjpvs",
	"https://imgur.com/EDWuNq8",
	"https://imgur.com/Eh33AVj",
	"https://imgur.com/FIwYIUQ",
	"https://imgur.com/FrmBadl",
	"https://imgur.com/HKxVnxG",
	"https://imgur.com/J7hHxO1",
	"https://imgur.com/L1N70Un",
	"https://imgur.com/LE4YpeZ",
	"https://imgur.com/LOXHumW",
	"https://imgur.com/O03Xja2",
	"https://imgur.com/QRLgUYx",
	"https://imgur.com/QfvSxYs",
	"https://imgur.com/TKUHvqm",
	"https://imgur.com/U6XRtLW",
	"https://imgur.com/UVc6OPm",
	"https://imgur.com/V86JYqi",
	"https://imgur.com/XZIe2uK",
	"https://imgur.com/a/3JLGzxc",
	"https://imgur.com/a/L7gzXp6",
	"https://imgur.com/a/LXR4tti",
	"https://imgur.com/a/ML0y9jf",
	"https://imgur.com/a/OKRwC4B",
	"https://imgur.com/a/S0aPnO9",
	"https://imgur.com/a/T1EsFz5",
	"https://imgur.com/a/Tqn1d90",
	"https://imgur.com/a/VT83R7p",
	"https://imgur.com/a/WQrkc8G",
	"https://imgur.com/a/cjUxAm3",
	"https://imgur.com/a/fDkOZfP",
	"https://imgur.com/a/x8Lo4Xj",
	"https://imgur.com/a/yPcmwaw",
	"https://imgur.com/a/zmKsRWX",
	"https://imgur.com/bIIlmlR",
	"https://imgur.com/d6CaKoz",
	"https://imgur.com/dYR2wMc",
	"https://imgur.com/h9ZOvcB",
	"https://imgur.com/hZKskr9",
	"https://imgur.com/iJFZgLf",
	"https://imgur.com/jv8vtTU",
	"https://imgur.com/muBRskQ",
	"https://imgur.com/nJjE0Pm",
	"https://imgur.com/o1nHlMW",
	"https://imgur.com/pAtUbQD",
	"https://imgur.com/pDtq7v0",
	"https://imgur.com/pMFcMlO",
	"https://imgur.com/skHLZSZ",
	"https://imgur.com/sxd3YPu",
	"https://imgur.com/tDfIoge",
	"https://imgur.com/tXHy8uu",
	"https://imgur.com/tsQz7ru",
	"https://imgur.com/vH2sy2N",
	"https://imgur.com/vUTd2zo",
	"https://imgur.com/wZILJLP",
	"https://lastfm-img2.akamaized.net/i/u/300x300/9d9f294f757a4ab782ca09078ee5ee1c.jpg",
	"https://lh3.googleusercontent.com/-bofEp1dLNgw/Wfkjm0C87wI/AAAAAAAAAko/nrPFiK35NIc_xeluS83WT-U4F78xhhshgCJoC/w1280-h1829/457f84e6-cadc-49b5-98e2-4f32705f11fd.png",
	"https://lh3.googleusercontent.com/-jekIwAo4P6c/XCK2RNE1BVI/AAAAAAAC4YM/31rmK6lpVhwMcaYasI4kViiV8_cuALpmACJoC/w1392-h1560-n-rw/vatvyr.png",
	"https://media1.tenor.com/images/93a28fa974e6e002f15fd5245a1079b7/tenor.gif?itemid=11578020",
	"https://memestatic.fjcdn.com/large/pictures/39/82/398288_6490166.jpg",
	"https://myswordisunbelievablydull.files.wordpress.com/2011/07/coalgirls_serial_experiments_lain_01_1008x720_blu-ray_flac_f0ef8af8-mkv_snapshot_16-02_2011-07-18_14-56-28.jpg",
	"https://myswordisunbelievablydull.files.wordpress.com/2011/07/coalgirls_serial_experiments_lain_03_1008x720_blu-ray_flac_92704257-mkv_snapshot_17-56_2011-07-19_19-09-07.jpg",
	"https://mywaifulist.s3-us-west-1.amazonaws.com/series/70/70.jpg",
	"https://nefariousreviews.files.wordpress.com/2016/09/serial-experiments-lain-obsession.jpg",
	"https://orig00.deviantart.net/041f/f/2016/273/d/b/lain_by_bronickelodeon-dajgd65.png",
	"https://outtherecinema.files.wordpress.com/2015/10/evil-lain.png",
	"https://pbs.twimg.com/media/Dv_13_XWwAEZaEK.jpg",
	"https://pbs.twimg.com/media/Dv_26MuW0AApOMD.jpg",
	"https://pbs.twimg.com/media/Dv_y_rkWsAAdCGs.jpg",
	"https://pbs.twimg.com/media/DwAm6JdW0AI-3Iz.jpg",
	"https://pbs.twimg.com/media/DwAmVyMX0AATe1e.jpg",
	"https://pbs.twimg.com/media/DwCLtWcWoAAy4jZ.jpg",
	"https://pbs.twimg.com/media/DwDnk1PUUAAGqER.jpg",
	"https://pbs.twimg.com/media/DwJx4LNVYAAbdBj.jpg",
	"https://pbs.twimg.com/media/DwK2TIkXcAAs6FJ.jpg",
	"https://pbs.twimg.com/media/DwKHX6QUYAEakTP.jpg",
	"https://pbs.twimg.com/media/Dx_D21qV4AAmkYT.jpg:large",
	"https://pm1.narvii.com/6194/ee5c2618687bbd1edd1afdc54f4e55309b7b80c6_hq.jpg",
	"https://pm1.narvii.com/6505/062a5a2d35200c8fa9916293025f484a6f3c57a4_hq.jpg",
	"https://raikou2.donmai.us/7f/2d/7f2d8cc8fc14b08373089a95c0a3182d.jpg",
	"https://s3.amazonaws.com/cult-of-distraction/media/articles/images/lain-13.png",
	"https://scontent-amt2-1.cdninstagram.com/vp/3aaab5cdf84dcba75ea9546789340e52/5CD0568C/t51.2885-15/sh0.08/e35/s640x640/42068850_620156955045839_471048183435590991_n.jpg?_nc_ht=scontent-amt2-1.cdninstagram.com",
	"https://scontent-atl3-1.cdninstagram.com/vp/61a380b854e01d4b1bb5a6b816beeb15/5C9B9A1B/t51.2885-15/e35/29093379_385671531907205_2690296831937609728_n.jpg",
	"https://scontent-sea1-1.cdninstagram.com/vp/c522fb17c6ed140811ec8b712968ce84/5CC7B08F/t51.2885-15/e35/c119.0.481.481/s480x480/28158331_1090161284476580_3726883881702391808_n.jpg?_nc_ht=scontent-sea1-1.cdninstagram.com",
	"https://static.zerochan.net/Iwakura.Lain.full.1207948.jpg",
	"https://static.zerochan.net/Iwakura.Lain.full.132356.jpg",
	"https://static.zerochan.net/Iwakura.Lain.full.1394713.jpg",
	"https://static.zerochan.net/Iwakura.Lain.full.485427.jpg",
	"https://static.zerochan.net/Iwakura.Lain.full.613356.jpg",
	"https://static.zerochan.net/Iwakura.Lain.full.769784.jpg",
	"https://static.zerochan.net/Iwakura.Lain.full.79038.jpg",
	"https://static.zerochan.net/Iwakura.Lain.full.891812.jpg",
	"https://steamusercontent-a.akamaihd.net/ugc/956347227393432230/4A37EBA7255E13D36505913E07D8BDCF94D6716D/",
	"https://steemitimages.com/640x0/https://i.hizliresim.com/1J37ZA.jpg",
	"https://stmed.net/sites/default/files/serial-experiments-lain-wallpapers-26100-9616498.jpg",
	"https://tacchi-canvas-prod.s3.amazonaws.com/uploads/work/image/5775/Serial_Experiments_Lain_by_dilara_ozden__Print_Designer_in_Tokyo.png",
	"https://thumbs.worthpoint.com/wpimages/images/images1/1/0917/23/1_867ebf074e5ad5d8180f0ba2f0c6f6f4.jpg",
	"https://upload.wikimedia.org/wikipedia/en/thumb/b/b8/Lain_hacker_small.jpg/250px-Lain_hacker_small.jpg",
	"https://vignette.wikia.nocookie.net/inciclopedia/images/1/16/Lain_oso.jpg/revision/latest?cb=20111226044923",
	"https://vignette.wikia.nocookie.net/malice-in-wonderland-and-all-things-alice/images/6/6c/Serial_experiments_lain_anime_desktop_wallpaper-normal.jpg/revision/latest?cb=20140214164314",
	"https://vignette.wikia.nocookie.net/sel/images/e/e4/Touko%27s_NAVI_Screen.png/revision/latest?cb=20100515114244",
	"https://wallup.net/wp-content/uploads/2016/05/14/52934-Serial_Experiments_Lain-Lain_Iwakura-748x421.png",
	"https://www.anbient.com/files/img/serial-experiments-lain/ss-8446.jpg",
	"https://www.downloadwallpapers.info/dl/1280x960//2016/02/21/21974_blood-girl-iwakura-lain-serial-experiments-lain-wall_1280x1024_h.jpg",
	"https://www.vizzed.com/videogames/psx/screenshot/Serial%20Experiments%20Lain-2.jpg",
	"https://www.wallpaperup.com/uploads/wallpapers/2014/03/15/298989/bf83071e887017b092582842dc759663-700.jpg",
	"https://zerojustice315.files.wordpress.com/2015/04/coalgirls_serial_experiments_lain_03_1008x720_blu-ray_flac_92704257-mkv_snapshot_16-51_2015-04-06_20-05-28.jpg",
}