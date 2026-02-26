# Отчёт по заданию — Git (trp-1)

## 1. Клонирование репозитория

```bash
git clone git@github.com:kirill010106/trp-1.git .
```
```
Cloning into '.'...
warning: You appear to have cloned an empty repository.
```

---

## 2. Создание SSH-ключа

```bash
ssh-keygen -t ed25519 -C "trp-1 task key" -f "$env:USERPROFILE\.ssh\trp1_task_key"
```
```
Generating public/private ed25519 key pair.
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in C:\Users\konst\.ssh\trp1_task_key
Your public key has been saved in C:\Users\konst\.ssh\trp1_task_key.pub
```

---

## 3. Создание файлов через консоль

```bash
Set-Content readme.txt "Project readme file"
Set-Content notes.txt "Project notes"
Set-Content config.txt "Project config"
```

---

## 4. Привязка локального репозитория к GitHub и первый коммит

```bash
git add .
git commit -m "Initial commit: add base files"
```
```
[master (root-commit) fbfe476] Initial commit: add base files
 3 files changed, 3 insertions(+)
 create mode 100644 config.txt
 create mode 100644 notes.txt
 create mode 100644 readme.txt
```

```bash
git push -u origin master
```
```
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 12 threads
Compressing objects: 100% (2/2), done.
Writing objects: 100% (5/5), 387 bytes | 129.00 KiB/s, done.
Total 5 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:kirill010106/trp-1.git
 * [new branch]      master -> master
branch 'master' set up to track 'origin/master'.
```

---

## 5. Создание новой ветки и вывод списка веток

```bash
git checkout -b feature
```
```
Switched to a new branch 'feature'
```

```bash
git branch -a
```
```
* feature
  master
  remotes/origin/main
  remotes/origin/master
```

---

## 6. Три коммита в новой ветке в разные файлы

```bash
Set-Content file1.txt "Feature data for file1"
git add file1.txt
git commit -m "feat: add file1"
```
```
[feature 575d6ed] feat: add file1
 1 file changed, 1 insertion(+)
 create mode 100644 file1.txt
```

```bash
Set-Content file2.txt "Feature data for file2"
git add file2.txt
git commit -m "feat: add file2"
```
```
[feature efacf9a] feat: add file2
 1 file changed, 1 insertion(+)
 create mode 100644 file2.txt
```

```bash
Set-Content file3.txt "Feature data for file3"
git add file3.txt
git commit -m "feat: add file3"
```
```
[feature 67de719] feat: add file3
 1 file changed, 1 insertion(+)
 create mode 100644 file3.txt
```

---

## 7. Выгрузка ветки в удалённый репозиторий

```bash
git push -u origin feature
```
```
Enumerating objects: 10, done.
Counting objects: 100% (10/10), done.
Delta compression using up to 12 threads
Compressing objects: 100% (6/6), done.
Writing objects: 100% (9/9), 927 bytes | 231.00 KiB/s, done.
Total 9 (delta 2), reused 0 (delta 0), pack-reused 0
To github.com:kirill010106/trp-1.git
 * [new branch]      feature -> feature
branch 'feature' set up to track 'origin/feature'.
```

---

## 8. Изменение файла без коммита + amend

```bash
Add-Content file3.txt "Extra line added before amend"
git status
```
```
On branch feature
Your branch is up to date with 'origin/feature'.

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   file3.txt

no changes added to commit (use "git add" and/or "git commit -a")
```

```bash
git add file3.txt
git commit --amend --no-edit
```
```
[feature 5802f5f] feat: add file3
 Date: Thu Feb 26 18:47:53 2026 +0300
 1 file changed, 2 insertions(+)
 create mode 100644 file3.txt
```

```bash
git log --oneline -5
```
```
5802f5f (HEAD -> feature) feat: add file3
efacf9a feat: add file2
575d6ed feat: add file1
fbfe476 (origin/master, origin/main, master) Initial commit: add base files
```

---

## 9. Различия между ветками master и feature

```bash
git diff master feature
```
```diff
diff --git a/file1.txt b/file1.txt
new file mode 100644
index 0000000..ac190ed
--- /dev/null
+++ b/file1.txt
@@ -0,0 +1 @@
+Feature data for file1
diff --git a/file2.txt b/file2.txt
new file mode 100644
index 0000000..13f435c
--- /dev/null
+++ b/file2.txt
@@ -0,0 +1 @@
+Feature data for file2
diff --git a/file3.txt b/file3.txt
new file mode 100644
index 0000000..1b3f8a1
--- /dev/null
+++ b/file3.txt
@@ -0,0 +1,2 @@
+Feature data for file3
+Extra line added before amend
```

---

## 10. Слияние feature с master

```bash
git checkout master
git merge feature
```
```
Switched to branch 'master'
Your branch is up to date with 'origin/master'.
Updating fbfe476..5802f5f
Fast-forward
 file1.txt | 1 +
 file2.txt | 1 +
 file3.txt | 2 ++
 3 files changed, 4 insertions(+)
 create mode 100644 file1.txt
 create mode 100644 file2.txt
 create mode 100644 file3.txt
```

```bash
git log --oneline -6
```
```
5802f5f (HEAD -> master, feature) feat: add file3
efacf9a feat: add file2
575d6ed feat: add file1
fbfe476 (origin/master, origin/main) Initial commit: add base files
```

```bash
git push origin master
```
```
Enumerating objects: 9, done.
Counting objects: 100% (9/9), done.
Delta compression using up to 12 threads
Compressing objects: 100% (5/5), done.
Writing objects: 100% (8/8), 821 bytes | 274.00 KiB/s, done.
Total 8 (delta 1), reused 0 (delta 0), pack-reused 0
To github.com:kirill010106/trp-1.git
   fbfe476..5802f5f  master -> master
```
