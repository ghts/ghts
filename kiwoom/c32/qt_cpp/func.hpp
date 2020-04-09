#ifndef FUNC_HPP
#define FUNC_HPP

#include <QWidget>

#define OK_ERR(condition) (condition ? " OK." : " Error.");

QString OkErr(bool condition);
bool LoadDLL();
bool InitGoDLL(QWidget *parent);
void NotifyLoginResult(bool loginResult);

#endif // FUNC_HPP
