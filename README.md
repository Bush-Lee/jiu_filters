# Ghid instalare

## Pasul 1 - Instalare git
Pentru a instala Git pe Windows, urmati acesti pasi:

1. Accesati site-ul oficial Git: https://git-scm.com/downloads si downloadati versiunea corespunzatoare sistemului de operare.
2. Descarcati fisierul de instalare pentru Windows.
3. Dupa ce ati descarcat fisierul, dublu clic pe acesta pentru a incepe procesul de instalare.
4. In fereastra de instalare, faceti clic pe "Next" pentru a continua. Puteti folosi setarile implicite pentru majoritatea optiunilor si faceti clic pe "Next".
5. Asigurati-va ca alegeti optiunea "Use Git from the Windows Command Prompt" cand apare, pentru a folosi Git din linia de comanda a Windows.
6. Asteptati ca instalarea sa se finalize si faceti clic pe "Finish" pentru a incheia procesul.

Dupa ce ati instalat Git, puteti deschide linia de comanda si verificati daca Git a fost instalat cu succes tastand git --version in fereastra command prompt.

## Pasul 2 - Descarcare proiect

1. Deschideti linia de comanda (Command Prompt sau PowerShell) pe computerul dumneavoastra.
2. Navigati catre directorul in care doriti sa clonati proiectul. Puteti folosi comanda cd cale_catre_directoriu pentru a naviga la directorul dorit. De exemplu: cd locatie\proiect\dorita
3. Pe pagina proiectului GitHub (https://github.com/Bush-Lee/jiu_filters), faceti clic pe butonul "Code" (situat in partea dreapta sus a paginii) pentru a obtine URL-ul de clonare a proiectului. (vezi poza x)
4. Copiati URL-ul proiectului. Puteti face acest lucru facand clic pe pictograma de copiere sau selectand manual URL-ul si apasand Ctrl + C.
5. Inapoi in linia de comanda, utilizati comanda git clone urmata de URL-ul pe care l-ati copiat. De exemplu: git clone https://github.com/Bush-Lee/jiu_filters.git
6. Apasati Enter pentru a executa comanda. Git va incepe procesul de clonare a proiectului de pe GitHub in directorul curent.
7. Asteptati pana cand Git finalizeaza clonarea proiectului. Vetii vedea progresul clonarii si o notificare atunci cand procesul este complet.

## Pasul 3 - Utilizare proiect 
1. Deschide PowerShell.
2. Navighează la directorul în care se află programul folosind comanda cd (Change Directory). De exemplu cd ```D:\jiu_filters```
3. După ce te afli în directorul corect, rulează programul folosind comanda ``` .\goServer\dataServer.exe run -d``` pentru a genera fisierele necesare crearii filtrelor.
4. Pentru a verifica rezultatele, ruleaza comanda ``` .\goServer\dataServer.exe run``` dupa care deschide fisierul ```main.html``` din folderul ```local``` 
5. Adauga modificarile: Foloseste git add pentru a adauga fisierele modificate sau noi pentru a fi incluse. De exemplu, ```git add .```
6. Comite modificările: Rulează ```git commit -m "Descriere modificari"``` pentru a crea un commit cu modificările pregătite.
7. Urca modficarile: Rulează ```git push ``` 
8. Autentificare (dacă este necesar): În unele cazuri (pasii 6, 7), vei fi solicitat să te autentifici pentru a trimite modificările. Urmează instrucțiunile de autentificare afișate în consolă.
