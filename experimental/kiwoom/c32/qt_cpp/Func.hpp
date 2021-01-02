#ifndef FUNC_HPP
#define FUNC_HPP

#include <QLibrary>
#include <QWidget>

#define OK_ERR(condition) (condition ? " OK." : " Error.");

bool LoadDLL();
QLibrary *GetKiwoom_Go();
bool Init(QWidget *parent);
void Confirm(uint serialNo, QString qString);

#endif // FUNC_HPP
