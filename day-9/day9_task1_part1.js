// 1a. Insert 5 operators pada table operators.
db.operators.insertMany([
    { "_id": 1, "name": "Telkomsel", "created_at": new Date(), "updated_at": new Date() },
    { "_id": 2, "name": "Smartfren", "created_at": new Date(), "updated_at": new Date() },
    { "_id": 3, "name": "XL", "created_at": new Date(), "updated_at": new Date() },
    { "_id": 4, "name": "Indosat", "created_at": new Date(), "updated_at": new Date() },
    { "_id": 5, "name": "Hinet", "created_at": new Date(), "updated_at": new Date() },
])
// 1b. Insert 3 product type.
db.product_types.insertMany([
    { "_id": 1, "name": "Pascabayar", "created_at": new Date(), "updated_at": new Date() },
    { "_id": 2, "name": "Prabayar", "created_at": new Date(), "updated_at": new Date() },
    { "_id": 3, "name": "Digital", "created_at": new Date(), "updated_at": new Date() },
])
// 1c. Insert 2 product dengan product type id = 1, dan operators id = 3.
db.products.insertMany([
    { "_id": 1, "name": "Silver", "product_type_id": 1, "operator_id": 3, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 2, "name": "Gold", "product_type_id": 1, "operator_id": 3, "created_at": new Date(), "updated_at": new Date() },
])
// 1d. Insert 3 product dengan product type id = 2, dan operator id = 1
db.products.insertMany([
    { "_id": 3, "name": "simPATI", "product_type_id": 2, "operator_id": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 4, "name": "KARTU As", "product_type_id": 2, "operator_id": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 5, "name": "LOOP", "product_type_id": 2, "operator_id": 1, "created_at": new Date(), "updated_at": new Date() },
])
// 1e. Insert 3 product dengan product type id = 3, dan operator id = 4
db.products.insertMany([
    { "_id": 6, "name": "Freedom Internet", "product_type_id": 3, "operator_id": 4, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 7, "name": "Freedom U As", "product_type_id": 3, "operator_id": 4, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 8, "name": "Freedom Combo", "product_type_id": 3, "operator_id": 4, "created_at": new Date(), "updated_at": new Date() },
])
// 1f. Insert product description pada setiap product.
db.product_descriptions.insertMany([
    { "_id": 1, "description": "Unlimited Nelp ke XL & AXIS\n100 Menit Nelp ke Operator Lain\nUnlimited SMS AnyNet™", "created_at": new Date(), "updated_at": new Date() },
    { "_id": 2, "description": "Unlimited Nelp ke XL & AXIS\n150 Menit Nelp ke Operator Lain\nUnlimited SMS AnyNet™, Total FUP 40GB", "created_at": new Date(), "updated_at": new Date() },
    { "_id": 3, "description": "Simpati (nama digayakan sebagai simPATI) adalah produk dari Telkomsel. Kartu ini diluncurkan pada tahun 1997. Pada tanggal 18 Juni 2021, kartu ini dileburkan bersama Kartu As dan LOOP menjadi Telkomsel PraBayar dengan diluncurkannya logo baru Telkomsel.", "created_at": new Date(), "updated_at": new Date() },
    { "_id": 4, "description": "Kartu As adalah produk dari Telkomsel. Kartu ini diluncurkan pada tahun 2004. Kartu As ini memiliki slogan 'Pas Buat Semua'. Pada tanggal 18 Juni 2021, kartu ini dileburkan bersama simPATI dan LOOP menjadi Telkomsel PraBayar dengan diluncurkannya logo baru Telkomsel.", "created_at": new Date(), "updated_at": new Date() },
    { "_id": 5, "description": "Telkomsel LOOP sebelumnya bernama simPATI LOOP atau yang dikenali sebagai LOOP adalah kartu GSM prabayar dari Telkomsel yang dirancang untuk kawula muda. Slogan yang dimiliki oleh LOOP adalah 'Ini Kita!' Pada tanggal 18 Juni 2021, kartu ini dileburkan bersama simPATI dan Kartu As menjadi Telkomsel PraBayar dengan diluncurkannya logo baru Telkomsel.", "created_at": new Date(), "updated_at": new Date() },
    { "_id": 6, "description": "Nikmati internetan 24 jam non-stop dengan kuota utama plus bebas nelpon sepuasnya ke sesama IM3 dan Tri selama 5000 menit. Saatnya kamu perluas jangkauan biar bisa ngobrol sama teman kapan pun di mana pun dan makin bebas buat ekspresikan lifestyle kamu lewat konten, sampai nonton YouTube atau main game sepuasnya pakai Freedom Internet dari IM3. Kapan lagi internetan dan nelpon bisa sehemat dan sepuas ini? Yuk, aktifkan segera Freedom Internet di *123#, melalui Aplikasi myIM3, atau IM3 Official WhatsApp!", "created_at": new Date(), "updated_at": new Date() },
    { "_id": 7, "description": "Sekarang kamu bisa makin puas #TerusTerusan akses aplikasi pakai Freedom U karena ada aplikasi-aplikasi tambahan baru seperti Netflix, Snapchat, Webex, Microsoft Teams, Skype, dan Google Classroom. Plus 24 jam bisa akses aplikasi lebih banyak dengan aplikasi favorit lainnya seperti YouTube, Instagram, TikTok, Facebook, Spotify, Joox, WhatsApp dan Line. Kamu juga bisa pakai Freedom U dengan aplikasi favorit lainnya seperti YouTube, Instagram, TikTok, Facebook, Spotify, Joox, WhatsApp dan Line. Tunggu apa lagi? Aktifkan Freedom U di *123#, myIM3, atau IM3 Ooredoo Official WhatsApp supaya makin puas internetan #TerusTerusan.", "created_at": new Date(), "updated_at": new Date() },
    { "_id": 8, "description": "Saatnya jalin komunikasi dengan yang terdekat pakai Freedom Combo! Dapatkan kuota besar hingga 60GB dan nelpon sepuasnya ke IM3 Ooredoo dan Tri. Nyaman non-stop telponan dan internetan dengan sahabat, teman dan kerabatmu tetap lancar di mana aja, kapan aja! Chatting, video call jadi bisa sepuasnya, nelpon juga lebih leluasa. Aktifkan di *123#, myIM3, atau IM3 Ooredoo Official WhatsApp.", "created_at": new Date(), "updated_at": new Date() },
])
// 1g. Insert 3 payment methods.
db.payment_methods.insertMany([
    { "_id": 1, "name": "Debit/Kredit", "created_at": new Date(), "updated_at": new Date() },
    { "_id": 2, "name": "GoPay", "created_at": new Date(), "updated_at": new Date() },
    { "_id": 3, "name": "OVO", "created_at": new Date(), "updated_at": new Date() },
])
// 1h. Insert 5 user pada tabel user.
db.users.insertMany([
    { "_id": 1, "name": 'Shiyam Lutfiyah Samaha', "dob": new Date('1959-03-17'), "status": 'KAWIN TERCATAT', "gender": 'PEREMPUAN', "created_at": new Date(), "updated_at": new Date() },
    { "_id": 2, "name": 'Hadi Khaldun Koury', "dob": new Date('1949-08-12'), "status": 'CERAI MATI', "gender": 'LAKI-LAKI', "created_at": new Date(), "updated_at": new Date() },
    { "_id": 3, "name": 'Elijah Farnell', "dob": new Date('1992-04-13'), "status": 'CERAI HIDUP', "gender": 'LAKI-LAKI', "created_at": new Date(), "updated_at": new Date() },
    { "_id": 4, "name": 'Brodie Hull', "dob": new Date('1993-11-7'), "status": 'KAWIN BELUM TERCATAT', "gender": 'LAKI-LAKI', "created_at": new Date(), "updated_at": new Date() },
    { "_id": 5, "name": 'Ashley Divine', "dob": new Date('2000-06-07'), "status": 'BELUM KAWIN', "gender": 'PEREMPUAN', "created_at": new Date(), "updated_at": new Date() },
])
// 1i. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1j)
db.transactions.insertMany([
    { "_id": 1, "user_id": 1, "total_price": 112000, "total_qty": 3, "payment_method_id": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 2, "user_id": 1, "total_price": 117000, "total_qty": 3, "payment_method_id": 2, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 3, "user_id": 1, "total_price": 60000, "total_qty": 3, "payment_method_id": 3, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 4, "user_id": 2, "total_price": 112000, "total_qty": 3, "payment_method_id": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 5, "user_id": 2, "total_price": 117000, "total_qty": 3, "payment_method_id": 2, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 6, "user_id": 2, "total_price": 60000, "total_qty": 3, "payment_method_id": 3, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 7, "user_id": 3, "total_price": 112000, "total_qty": 3, "payment_method_id": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 8, "user_id": 3, "total_price": 117000, "total_qty": 3, "payment_method_id": 2, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 9, "user_id": 3, "total_price": 60000, "total_qty": 3, "payment_method_id": 3, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 10, "user_id": 4, "total_price": 112000, "total_qty": 3, "payment_method_id": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 11, "user_id": 4, "total_price": 117000, "total_qty": 3, "payment_method_id": 2, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 12, "user_id": 4, "total_price": 60000, "total_qty": 3, "payment_method_id": 3, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 13, "user_id": 5, "total_price": 112000, "total_qty": 3, "payment_method_id": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 14, "user_id": 5, "total_price": 117000, "total_qty": 3, "payment_method_id": 2, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 15, "user_id": 5, "total_price": 60000, "total_qty": 3, "payment_method_id": 3, "created_at": new Date(), "updated_at": new Date() },
])
// 1j. Insert 3 product di masing-masing transaksi.
db.transaction_details.insertMany([
    { "_id": 1, "transaction_id": 1, "product_id": 8, "price": 25000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 2, "transaction_id": 1, "product_id": 1, "price": 80000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 3, "transaction_id": 1, "product_id": 6, "price": 7000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 4, "transaction_id": 2, "product_id": 7, "price": 5000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 5, "transaction_id": 2, "product_id": 2, "price": 102000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 6, "transaction_id": 2, "product_id": 5, "price": 10000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 7, "transaction_id": 3, "product_id": 2, "price": 30000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 8, "transaction_id": 3, "product_id": 4, "price": 20000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 9, "transaction_id": 3, "product_id": 5, "price": 10000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 10, "transaction_id": 4, "product_id": 8, "price": 25000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 11, "transaction_id": 4, "product_id": 1, "price": 80000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 12, "transaction_id": 4, "product_id": 6, "price": 7000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 13, "transaction_id": 5, "product_id": 7, "price": 5000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 14, "transaction_id": 5, "product_id": 2, "price": 102000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 15, "transaction_id": 5, "product_id": 5, "price": 10000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 16, "transaction_id": 6, "product_id": 2, "price": 30000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 17, "transaction_id": 6, "product_id": 4, "price": 20000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 18, "transaction_id": 6, "product_id": 5, "price": 10000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 19, "transaction_id": 7, "product_id": 8, "price": 25000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 20, "transaction_id": 7, "product_id": 1, "price": 80000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 21, "transaction_id": 7, "product_id": 6, "price": 7000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 22, "transaction_id": 8, "product_id": 7, "price": 5000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 23, "transaction_id": 8, "product_id": 2, "price": 102000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 24, "transaction_id": 8, "product_id": 5, "price": 10000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 25, "transaction_id": 9, "product_id": 2, "price": 30000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 26, "transaction_id": 9, "product_id": 4, "price": 20000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 27, "transaction_id": 9, "product_id": 5, "price": 10000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 28, "transaction_id": 10, "product_id": 8, "price": 25000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 29, "transaction_id": 10, "product_id": 1, "price": 80000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 30, "transaction_id": 10, "product_id": 6, "price": 7000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 31, "transaction_id": 11, "product_id": 7, "price": 5000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 32, "transaction_id": 11, "product_id": 2, "price": 102000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 33, "transaction_id": 11, "product_id": 5, "price": 10000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 34, "transaction_id": 12, "product_id": 2, "price": 30000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 35, "transaction_id": 12, "product_id": 4, "price": 20000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 36, "transaction_id": 12, "product_id": 5, "price": 10000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 37, "transaction_id": 13, "product_id": 8, "price": 25000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 38, "transaction_id": 13, "product_id": 1, "price": 80000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 39, "transaction_id": 13, "product_id": 6, "price": 7000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 40, "transaction_id": 14, "product_id": 7, "price": 5000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 41, "transaction_id": 14, "product_id": 2, "price": 102000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 42, "transaction_id": 14, "product_id": 5, "price": 10000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 43, "transaction_id": 15, "product_id": 2, "price": 30000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 44, "transaction_id": 15, "product_id": 4, "price": 20000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
    { "_id": 45, "transaction_id": 15, "product_id": 5, "price": 10000, "qty": 1, "created_at": new Date(), "updated_at": new Date() },
])
// 2a. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
db.users.find({ "gender": "LAKI-LAKI" })
// 2b. Tampilkan product dengan id = 3.
db.products.find({ _id: 3 })
// 2c. Hitung jumlah user / pelanggan dengan status gender Perempuan.
db.users.aggregate([
    { $match: { "gender": "PEREMPUAN" }, },
    { $count: "user_perempuan", },
])
// 2d. Tampilkan data pelanggan dengan urutan sesuai nama abjad.
db.users.find().sort({ "name": 1 })
// 2e. Tampilkan 5 data product.
db.products.find().limit(2)
// 3a. Ubah data product id 1 dengan nama 'product dummy'.
db.products.updateOne(
    { "_id": 1 },
    { $set: { "name": "product_dummy" } },
)
// 3b. Ubah qty = 3 pada transaction detail dengan product id 1.
db.transaction_details.updateMany(
    { "product_id": 1 },
    { $set: { "qty": 3 } }
)
// 4a. Delete data pada tabel product dengan id 1.
db.products.deleteOne({ "_id": 1 })
// 4b. Delete pada pada tabel product dengan product type id 1.
db.products.deleteMany({ "product_type_id": 1 })