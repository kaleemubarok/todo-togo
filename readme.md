**Final Project**
Buatlah sebuah APP service (web-server) dimana service "ToDo Services" yang akan dipublish dan release ke public sebagai suatu layanan Catatan Online. Fitur-fitur yang dimiliki oleh APP ini adalah
1. Fitur
    - Membuat sebuah ToDo Task (POST) 
        - Status ToDo : New, OnGoing, Done, Deleted
    - Mengedit sebuah ToDo Task (PUT) 
        - Jika ToDo dalam status Done atau Deleted, maka dia tidak dapat diubah
        - Editing dapat dilakukan ke semua informasi ToDo
    - Mendelete sebuah ToDo Task (DELETE) 
    - Menampilkan seluruh ToDo Task (GET) 
    - Menambah User (POST) 
    - Mengedit User (PUT) 
    - Mendelete User (DELETE) 
    
    You may refer:
    ```
    Sesi 2,3,6,8,11
2. Object
    1. ToDo
        - Title (text)
        - Description (text)
        - Due Date (date)
        - Person In Charge (string)
        - Status (string)
    2. User
        - UserID (int)
        - Name (text/string)
    3. Status
        - StatusID (int)
        - StatusTxt (string)
***
Criteria :
- Implement *API*. Responsenya berupa json
- Implement *Swaggo Documentation*
- *Unit Test* minimal 50%
- Semua yang dijelaskan adalah *WAJIB*, selain yang dijelaskan kalian bebas untuk membuat asumsi/style masing-masing, misalkan library, framework, code structure, etc
Push assignment ke github masing-masing.
Buat file txt dengan notepad atau editor yang kalian miliki, masukan link repository assignment kalian kedalam file txt tersebut. Unggah file .txt tersebut dalam kotak _*Submit Your Assignment*_ di *SESI12*.
Submission Assignment *Valid* jika link yang kalian submit dapat diakses kemudian dapat dinilai oleh Hacktiv8 PTP Program Code Reviewer.